package api

import (
	"encoding/json"
	"net/http"
)

type LocationResponse struct {
	Count        int     `json:"count"`
	Next         *string `json:"next"`
	Previous     *string `json:"previous"`
	LocationList []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

func (c *Client) FetchLocations(endpoint *string) (LocationResponse, error) {
	url := base_url
	if endpoint != nil {
		url = url + *endpoint
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}

	var locationsData LocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locationsData)
	if err != nil {
		return LocationResponse{}, err
	}

	return locationsData, nil
}
