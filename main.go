package main

import (
	"time"

	"github.com/statusquonjc46/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := config{
		commands:      getCommands(),
		pokeApiClient: pokeClient,
		pokemon:       map[string]pokeapi.Pokemon{},
	}
	startRepl(&cfg)
}
