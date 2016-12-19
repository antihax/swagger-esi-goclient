java -jar swagger-codegen-cli.jar generate -o ./tests/go-petstore -t ./template -l go -i ./tests/petstore-with-fake-endpoints-models-for-testing.yaml
go run tools/postgen.go -src ./tests/go-petstore
cd ./tests
go get
go test