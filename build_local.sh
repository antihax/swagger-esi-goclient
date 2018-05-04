#!/bin/bash

go get github.com/sqs/goreturns
go get -u github.com/mailru/easyjson/...

rm -rf ../goesi/esi/*
rm -rf ../goesi/esi/docs/*
set -e
# Generate models first and JSON code second because there is no easy way to glob for model files only
echo Build models
java -jar -Dmodels ../swagger-esi-goclient/swagger-codegen-cli.jar generate -o ../goesi/esi -t ../swagger-esi-goclient/template -l go -i https://esi.evetech.net/_latest/swagger.json?datasource=tranquility -DpackageName=esi
echo format models
find ../goesi/esi/ -type f -name "*.go" -exec echo processing {} \; -exec easyjson -noformat {} \;
# Generate all the other files
echo "regenerate"
java -jar ../swagger-esi-goclient/swagger-codegen-cli.jar generate -o ../goesi/esi -t ../swagger-esi-goclient/template -l go -i https://esi.evetech.net/_latest/swagger.json?datasource=tranquility -DpackageName=esi

echo fix slices of structs
# Fix slices of struct types
sed -i 's/REMOVEME\[\]//g' ../goesi/esi/*.*

echo fix imports and simplify
# Fix imports where needed (select encoding/json or easyjson)
goreturns -l -w ../goesi/esi

echo format code
set -e
gofmt -s -w ../goesi

sed -i 's/antihax\/optional/antihax\/goesi\/optional/g' ../goesi/esi/*.*

echo test
go test ./...


