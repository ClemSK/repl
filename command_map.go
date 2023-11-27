package main

import (
	"fmt"
	"log"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.nextLocationAreaUrl = resp.Previous
	return nil
}
