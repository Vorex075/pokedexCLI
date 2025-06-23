package commands

import (
	"bufio"
	"fmt"
	"os"
)

func commandInspect(cfg *Config, args []string) error {
	var pokemonName string
	if args == nil || len(args) == 0 {
		lineReader := bufio.NewScanner(os.Stdin)
		fmt.Printf("Introduce a pokemon name: ")
		if !lineReader.Scan() {
			return fmt.Errorf("error while trying to read from command line: %v",
				lineReader.Err())
		}
		pokemonName = lineReader.Text()
	} else {
		pokemonName = args[0]
	}

	_, err := cfg.pokeapiClient.FetchPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	pokeData, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n",
		pokemonName, pokeData.info.Height, pokeData.info.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokeData.info.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseValue)
	}
	fmt.Println("Types:")
	for _, typeValue := range pokeData.info.Types {
		fmt.Printf("\t- %s\n", typeValue.Type.Name)
	}
	return nil
}
