package commands

import (
	"github.com/Vorex075/pokedexCLI/internal/api"
)

type Config struct {
	pokeapiClient api.Client
	next          *string
	prev          *string
}
