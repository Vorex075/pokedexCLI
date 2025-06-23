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
	Types          []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseValue int `json:"base_stat"`
		Stat      struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Weight int `json:"weight"`
	Height int `json:"height"`
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
