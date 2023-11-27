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

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Print(err)
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
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Lists location areas",
			callback:    commandMap,
		},
		// "mapb": {
		// 	name:        "mapb",
		// 	description: "Displays the previous 20 location areas",
		// 	callback:    commandMapb,
		// },
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
