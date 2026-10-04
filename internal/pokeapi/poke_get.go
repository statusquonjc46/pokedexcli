package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationData := RespLocations{}
		err := json.Unmarshal(val, &locationData)
		if err != nil {
			return RespLocations{}, err
		}
		return locationData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationData := RespLocations{}
	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, data)
	return locationData, nil
}
