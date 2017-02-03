cls
COUNTER=1
while [  $COUNTER -lt 5 ]; do
    rm ../esi/v$COUNTER/*
    java -jar swagger-codegen-cli.jar generate -o ../esi/v$COUNTER -t ./template -l go -i https://esi.tech.ccp.is/v$COUNTER/swagger.json?datasource=tranquility -DpackageName=esiv$COUNTER
    let COUNTER=COUNTER+1 
done