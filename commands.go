package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, val := range cfg.commands {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}

func commandMapF(cfg *config) error {
	return nil
}

func commandMapB(cfg *config) error {
	return nil
}
