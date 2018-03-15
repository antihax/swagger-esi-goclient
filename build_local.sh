#!/bin/bash

rm -rf ../goesi/esi/*
rm -rf ../goesi/esi/docs/*
set -e
# Generate models first and JSON code second because there is no easy way to glob for model files only
java -jar -Dmodels ./swagger-codegen-cli.jar generate -o ../goesi/esi -t ./template -l go -i https://esi.evetech.net/_latest/swagger.json?datasource=tranquility -DpackageName=esi
find ../goesi/esi/ -type f -name "*.go" -exec echo processing {} \; -exec easyjson -noformat {} \;
# Generate all the other files
java -jar ./swagger-codegen-cli.jar generate -o ../goesi/esi -t ./template -l go -i https://esi.evetech.net/_latest/swagger.json?datasource=tranquility -DpackageName=esi
# Fix slices of struct types
sed -i 's/REMOVEME\[\]//g' ../goesi/esi/*.*
# Fix imports where needed (select encoding/json or easyjson)
goreturns -w ../goesi/esi
let COUNTER=COUNTER+1 
set +e

gofmt -s -w ../goesi
go test ../goesi/...