package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonInfo struct {
	Name           string `json:"name"`
	Id             int    `json:"id"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Client) FetchPokemonInfo(pokemonName string) (PokemonInfo, error) {
	url := base_url + "/pokemon/" + pokemonName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return PokemonInfo{}, fmt.Errorf("'%s' is not a valid pokemon name...",
			pokemonName)
	}
	var pokemonData PokemonInfo
	err = json.NewDecoder(resp.Body).Decode(&pokemonData)
	if err != nil {
		return PokemonInfo{}, err
	}
	return pokemonData, nil
}
