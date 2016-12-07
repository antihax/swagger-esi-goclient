# swagger-esi-goclient

swagger codegen template for EVE's ESI API

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
