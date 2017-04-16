#!/bin/bash

COUNTER=1

while [ $COUNTER -lt 5 ]; do
    rm ../goesi/v$COUNTER/*
    rm ../goesi/v$COUNTER/docs/*
    set -e
    java -jar ./swagger-codegen-cli.jar generate -o ../goesi/v$COUNTER -t ./template -l go -i https://esi.tech.ccp.is/v$COUNTER/swagger.json?datasource=tranquility -DpackageName=goesiv$COUNTER
    let COUNTER=COUNTER+1 
    set +e
done

gofmt -s -w ../goesi
go test ../goesi/...