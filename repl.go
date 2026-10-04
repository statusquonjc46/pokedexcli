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
	callback    func(*config, string) error
}

type config struct {
	commands      map[string]cliCommand
	pokeApiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
	pokemon       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	bufferScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanBool := bufferScanner.Scan()
		if scanBool == false {
			fmt.Println(bufferScanner.Err())
		}
		cleanedInput := cleanInput(bufferScanner.Text())
		argTwo := ""

		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]

		switch commandName {
		case "explore":
			argTwo = cleanedInput[1]
		case "catch":
			argTwo = cleanedInput[1]
		case "inspect":
			argTwo = cleanedInput[1]
		default:
			argTwo = ""
		}

		//fmt.Printf("Your command was: %s\n", commandName)

		cmd, exists := cfg.commands[commandName]
		if exists {
			err := cmd.callback(cfg, argTwo)
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
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Print all pokemon in the chosen area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the named pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Print a caught pokemon's stats.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print out the name for all pokemon caught.",
			callback:    commandPokedex,
		},
	}
}
