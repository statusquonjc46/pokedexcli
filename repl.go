package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/statusquonjc46/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands      map[string]cliCommand
	pokeApiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func startRepl(cfg *config) {
	bufferScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex >")
		bufferScanner.Scan()
		cleanedInput := cleanInput(bufferScanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]

		//fmt.Printf("Your command was: %s\n", commandName)

		cmd, exists := cfg.commands[commandName]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	//splitText := strings.Split(strings.Join(strings.Fields(strings.ToLower(text)), " "), " ")
	//splitText := strings.Fields(strings.ToLower(text))
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
	//return splitText
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints out the list of possible commands.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the first and next 20 locations.",
			callback:    commandMapF, //fix
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapB, //fix
		},
	}
}
