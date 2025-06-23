package commands

import (
	"fmt"
)

func commandMap(cfg *Config, _ []string) error {
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

func commandMapb(cfg *Config, _ []string) error {
	if cfg.prev == nil {
		fmt.Println("you are on the first page")
		return nil
	}
	locationResp, err := cfg.pokeapiClient.FetchLocations(cfg.prev)
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
