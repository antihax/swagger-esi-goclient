cls
rm ./esi/*
java -jar swagger-codegen-cli.jar generate -o ./esi -t ./template -l go -i https://esi.tech.ccp.is/latest/swagger.json?datasource=tranquility -DpackageName=esi