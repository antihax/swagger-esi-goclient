#!/bin/bash

rev=$(git rev-parse --short HEAD)

cd ..
mkdir goesi
cd goesi
git init
git config --global user.name "${GH_USER}"
git config --global user.email "${GH_EMAIL}"
git config --global github.token "${GH_TOKEN}"
git config --global push.default simple

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
git commit -m "rebuild pages at ${rev}"
set +e
git push -q upstream HEAD
set -e

go get -v
go test ./...