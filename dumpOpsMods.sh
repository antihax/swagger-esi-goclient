cls
java -DdebugModels -jar swagger-codegen-cli.jar generate -o ./dump -t ./template -l go -i https://esi.evetech.net/latest/swagger.json?datasource=tranquility > mods.txt
java -DdebugOperations -jar swagger-codegen-cli.jar generate -o ./dump -t ./template -l go -i https://esi.evetech.net/latest/swagger.json?datasource=tranquility > ops.txt