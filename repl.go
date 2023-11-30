package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) { // sharing access to a config struct to avoid copying state
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		cleaned := cleanInput(reader.Text())
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		command, exists := getCommands()[commandName]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if exists {
			err := command.callback(cfg, args...)
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
	output := strings.ToLower(text)
	cleaned := strings.Fields(output)
	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location areas",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous page of location areas",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
