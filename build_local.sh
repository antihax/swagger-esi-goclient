#!/bin/bash

COUNTER=1

while [ $COUNTER -lt 5 ]; do
    # Cleanup
    rm ../goesi/v$COUNTER/*
    rm ../goesi/v$COUNTER/docs/*
    set -e
    # Generate models first and JSON code second because there is no easy way to glob for model files only
    java -jar -Dmodels ./swagger-codegen-cli.jar generate -o ../goesi/v$COUNTER -t ./template -l go -i https://esi.tech.ccp.is/v$COUNTER/swagger.json?datasource=tranquility -DpackageName=goesiv$COUNTER
    easyjson -noformat ../goesi/v$COUNTER/*.go
    # Generate all the other files
    java -jar ./swagger-codegen-cli.jar generate -o ../goesi/v$COUNTER -t ./template -l go -i https://esi.tech.ccp.is/v$COUNTER/swagger.json?datasource=tranquility -DpackageName=goesiv$COUNTER
    # Fix slices of struct types
    sed -i 's/REMOVEME\[\]//g' ../goesi/v$COUNTER/*.*
    # Fix imports where needed (select encoding/json or easyjson)
    goreturns -w ../goesi/v$COUNTER
    let COUNTER=COUNTER+1 
    set +e
done

gofmt -s -w ../goesi
go test ../goesi/...