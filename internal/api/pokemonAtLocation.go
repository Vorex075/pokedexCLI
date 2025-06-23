package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonAtLocation struct {
	LocationName string `json:"name"`
	Pokemons     []struct {
		PokemonInfo struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetPokemonsAtLocation(locationName string) (PokemonAtLocation, error) {
	url := base_url + "/location-area/" + locationName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonAtLocation{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonAtLocation{}, err
	}
	if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
		return PokemonAtLocation{}, fmt.Errorf("it seems %s is not a valid city...",
			locationName)
	}
	var pokemonData PokemonAtLocation
	err = json.NewDecoder(resp.Body).Decode(&pokemonData)
	if err != nil {
		return PokemonAtLocation{}, err
	}
	return pokemonData, nil
}
