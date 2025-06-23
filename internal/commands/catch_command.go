package commands

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/Vorex075/pokedexCLI/internal/api"
)

type Pokemon struct {
	info      api.PokemonInfo
	catchTime time.Time
}

func CommandCatch(cfg *Config, args []string) error {
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

	data, err := cfg.pokeapiClient.FetchPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	isCatch := int(math.Log2(float64(data.BaseExperience))) % ((rand.Int() % 5) + 1)
	fmt.Printf("Is catch: %v\n", isCatch)
	if !(isCatch <= 3) {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	newCatch := Pokemon{
		info:      data,
		catchTime: time.Now(),
	}
	cfg.pokedex[pokemonName] = newCatch
	return nil
}
