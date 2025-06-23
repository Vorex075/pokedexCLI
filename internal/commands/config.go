package commands

import (
	"github.com/Vorex075/pokedexCLI/internal/api"
	"github.com/Vorex075/pokedexCLI/internal/pokecache"
)

type Config struct {
	pokeapiClient api.Client
	cache         *pokecache.Cache
	next          *string
	prev          *string
	pokedex       map[string]Pokemon
}

func NewConfig(client api.Client, cache *pokecache.Cache) Config {
	return Config{
		pokeapiClient: client,
		cache:         cache,
		pokedex:       make(map[string]Pokemon),
	}
}
