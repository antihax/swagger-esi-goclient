#!/bin/bash
go get -v
go get github.com/sqs/goreturns
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

echo fix imports and simplify
# Fix imports where needed (select encoding/json or easyjson)
goreturns -l -w ../goesi/esi
goreturns -l -w ../goesi/meta

echo format code
set -e
gofmt -s -w ../goesi

echo test
go test ./...


