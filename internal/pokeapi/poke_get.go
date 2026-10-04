package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return RespLocations{}, fmt.Errorf("Status Code Returned: %v", resp.StatusCode)
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

func (c *Client) GetPokemonInArea(area string) (AreaData, error) {
	if area == "" {
		return AreaData{}, errors.New("No Area Provided.")
	}
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		areaData := AreaData{}
		err := json.Unmarshal(val, &areaData)
		if err != nil {
			return AreaData{}, err
		}
		return areaData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaData{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return AreaData{}, fmt.Errorf("Status Code Returned: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaData{}, err
	}

	areaData := AreaData{}
	err = json.Unmarshal(data, &areaData)
	if err != nil {
		return AreaData{}, err
	}

	c.cache.Add(url, data)
	return areaData, nil
}

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	if pokemon == "" {
		return Pokemon{}, errors.New("No Pokemon Provided.")
	}
	url := baseURL + "/pokemon/" + pokemon
	//fmt.Printf("DEBUG: %v", url)

	if val, ok := c.cache.Get(url); ok {
		pokemonData := Pokemon{}
		err := json.Unmarshal(val, &pokemonData)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return Pokemon{}, fmt.Errorf("Status Code Returned: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonData, nil
}
