package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	fmt.Println("Location areas")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	fmt.Println("Location areas")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
