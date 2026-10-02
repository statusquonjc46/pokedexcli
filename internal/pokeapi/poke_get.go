package pokeapi

//figure out logic for http client, mapF, mapB, etc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getLocationAreas() ([]byte, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/{id or name}/")
	if err != nil {
		return nil, errors.New("Failed to query PokeAPI.")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Failed GET with status code: %v", res.StatusCode)
	}
	if err != nil {
		return nil, errors.New("Failed to read Body")
	}

	return body, nil
}
