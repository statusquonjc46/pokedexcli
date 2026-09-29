package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func printHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func startRepl() {
	bufferScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex >")
		bufferScanner.Scan()
		cleanedInput := cleanInput(bufferScanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]

		fmt.Printf("Your command was: %s\n", commandName)

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback()
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
	splitText := strings.Split(strings.Join(strings.Fields(strings.ToLower(text)), " "), " ")
	//splitText := strings.Fields(strings.ToLower(text))
	return splitText
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
			callback:    printHelp,
		},
	}
}
