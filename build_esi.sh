cls
COUNTER=1
while [  $COUNTER -lt 5 ]; do
    rm ../esi/v$COUNTER/*
    rm ../esi/v$COUNTER/doc/*
    java -jar swagger-codegen-cli.jar generate -o ../goesi/v$COUNTER -t ./template -l go -i https://esi.tech.ccp.is/v$COUNTER/swagger.json?datasource=tranquility -DpackageName=goesiv$COUNTER
    let COUNTER=COUNTER+1 
done