package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, _ string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, val := range cfg.commands {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}

func commandMapF(cfg *config, _ string) error {
	locations, err := cfg.pokeApiClient.GetLocationAreas(cfg.nextLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locations.Next
	cfg.prevLocation = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, _ string) error {
	locations, err := cfg.pokeApiClient.GetLocationAreas(cfg.prevLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locations.Next
	cfg.prevLocation = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config, area string) error {
	areaData, err := cfg.pokeApiClient.GetPokemonInArea(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", area)
	for _, p := range areaData.PokemonEncounters {
		currPokemon := p.Pokemon.Name
		fmt.Printf(" - %v\n", currPokemon)
	}

	return nil
}

func commandCatch(cfg *config, pokemon string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	//const maxBaseExp int = 650
	pokemonDetails, err := cfg.pokeApiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	//fmt.Println("DEBUG base experience:", pokemonDetails)
	randomChance := rand.Intn(pokemonDetails.BaseExperience)

	if randomChance < 40 {
		fmt.Printf("%v was caught!\n", pokemon)
		cfg.pokemon[pokemon] = pokemonDetails
	} else {
		fmt.Printf("%v escaped!\n", pokemon)
	}

	return nil
}
