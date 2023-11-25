package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

// func startRepl
func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
	}
}

// func cleanInput

// type cliStruct

// func getCommands

//
