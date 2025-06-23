package commands

import (
	"bufio"
	"fmt"
	"os"
)

func commandExplore(cfg *Config, args []string) error {
	var location string
	if args == nil || len(args) == 0 {
		lineReader := bufio.NewScanner(os.Stdin)
		fmt.Printf("Introduce a location: ")
		if !lineReader.Scan() {
			return fmt.Errorf("error while trying to read from command line: %v",
				lineReader.Err())
		}
		location = lineReader.Text()
	} else {
		location = args[0]
	}
	pokemonData, err := cfg.pokeapiClient.GetPokemonsAtLocation(location)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemons at '%s':\n", location)
	for _, pokemon := range pokemonData.Pokemons {
		fmt.Printf("\t- %s\n", pokemon.PokemonInfo.Name)
	}
	return nil
}
