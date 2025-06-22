package commands

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	locationResp, err := cfg.pokeapiClient.FetchLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locationResp.Next
	cfg.prev = locationResp.Previous

	for _, loc := range locationResp.LocationList {
		fmt.Println(loc.Name)
	}
	return nil
}
