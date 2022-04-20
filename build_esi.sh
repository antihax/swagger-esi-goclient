#!/bin/bash

rev=$(git rev-parse --short HEAD)

git config --global user.name "${GH_USER}"
git config --global user.email "${GH_EMAIL}"
git config --global github.token "${GH_TOKEN}"
git config --global push.default simple

cd ..
mkdir mock-esi
cd mock-esi
git init
git remote add upstream "https://${GH_TOKEN}@github.com/antihax/mock-esi.git"
git fetch upstream
git pull upstream HEAD
touch version.txt

oldnum=`cut -d ',' -f2 version.txt`  
newnum=`expr $oldnum + 1`
sed -i "s/$oldnum\$/$newnum/g" version.txt 
set -e
go get -v
bash build_mock_esi.sh
gofmt -s -w .
set +e
git add -A .
git commit -m "rebuild esi at ${rev}"
git push -q upstream HEAD

cd ..
mkdir goesi
cd goesi
git init
git remote add upstream "https://${GH_TOKEN}@github.com/antihax/goesi.git"
git fetch upstream
git pull upstream HEAD
go get -v
go get -u github.com/mailru/easyjson/...
go get github.com/antihax/optional

rm -rf ../goesi/esi/*
rm -rf ../goesi/esi/docs/*
rm -rf ../goesi/meta/*
rm -rf ../goesi/meta/docs/*

set +e
# Generate models first and JSON code second because there is no easy way to glob for model files only
echo Build models
java -jar -Dmodels ../swagger-esi-goclient/swagger-codegen-cli.jar generate -o ../goesi/esi -t ../swagger-esi-goclient/template -l go -i https://esi.evetech.net/_latest/swagger.json -DpackageName=esi
echo format models
find ../goesi/esi/ -type f -name "*.go" -exec echo processing {} \; -exec easyjson -noformat {} \;
# Generate all the other files
echo "regenerate"
java -jar ../swagger-esi-goclient/swagger-codegen-cli.jar generate -o ../goesi/esi -t ../swagger-esi-goclient/template -l go -i https://esi.evetech.net/_latest/swagger.json -DpackageName=esi

echo fix slices of structs ESI
# Fix slices of struct types
sed -i 's/REMOVEME\[\]//g' ../goesi/esi/*.*

# Generate models first and JSON code second because there is no easy way to glob for model files only
echo Build models
java -jar -Dmodels ../swagger-esi-goclient/swagger-codegen-cli.jar generate -o ../goesi/meta -t ../swagger-esi-goclient/template -l go -i https://esi.evetech.net/swagger.json -DpackageName=meta
echo format models
find ../goesi/meta/ -type f -name "*.go" -exec echo processing {} \; -exec easyjson -noformat {} \;
# Generate all the other files
echo "regenerate"
java -jar ../swagger-esi-goclient/swagger-codegen-cli.jar generate -o ../goesi/meta -t ../swagger-esi-goclient/template -l go -i https://esi.evetech.net/swagger.json -DpackageName=meta

sed -i 's/antihax\/optional/antihax\/goesi\/optional/g' ../goesi/esi/*.*
sed -i 's/antihax\/optional/antihax\/goesi\/optional/g' ../goesi/meta/*.*

echo fix slices of structs meta
# Fix slices of struct types
sed -i 's/REMOVEME\[\]//g' ../goesi/meta/*.*

echo format code
set -e
gofmt -s -w ../goesi

echo test
go test ./...
git add -A .

set +e
git commit -m "rebuild esi at ${rev}"
git push -q upstream HEAD
set -e

go get -v
go test ./...
