# swagger-esi-goclient

swagger codegen template for EVE's ESI API

## Sample Ouput
See sample output [https://github.com/antihax/evedata/tree/master/esi](https://github.com/antihax/evedata/tree/master/esi)

## Generating the code

The output directory can be specified with `-o`.

Run a command similar to the following to generate the code:

```java -jar swagger-codegen-cli.jar generate -o ../esi -t ./template -l go -i https://esi.tech.ccp.is/latest/swagger.json?datasource=tranquility -DpackageName=esi```

## Usage Examples

List Alliances:

    client := esi.NewAPIClient(nil)
    alliances, err := client.AllianceApi.GetAlliances("tranquility")
    if err != nil {
        log.Fatalf("Failed to get alliances: %v\n", err)
    }
    for _, allianceID := range alliances {
            log.Printf("\t Alliance ID: %d\n", allianceID)
        }

    log.Printf("====> There are %d alliances\n", len(alliances))

Get Solar System Info:

    client := esi.NewAPIClient(nil)
    systems, err := client.UniverseApi.PostUniverseNames(esi.PostUniverseNamesIds{
		Ids: []int32{30000142},
	}, "tranquility")

	if err != nil {
		log.Fatalf("Failed to lookup systems: %v\n", err)
	}
	fmt.Println(systems)
