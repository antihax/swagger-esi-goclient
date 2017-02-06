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

COUNTER=1
while [ $COUNTER -lt 5 ]; do
    rm ./v$COUNTER/*
    rm ./v$COUNTER/doc/*
    set -e
    java -jar ../swagger-esi-goclient/swagger-codegen-cli.jar generate -o ./v$COUNTER -t ../swagger-esi-goclient/template -l go -i https://esi.tech.ccp.is/v$COUNTER/swagger.json?datasource=tranquility -DpackageName=goesiv$COUNTER
    let COUNTER=COUNTER+1 
    set +e
done

set -e

go test ./...
git add -A .

set +e
git commit -m "rebuild esi at ${rev}"
git push -q upstream HEAD
set -e

go get -v
go test ./...