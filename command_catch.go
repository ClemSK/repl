package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	// we use base experience (int) to determine how hard a pokemon is to catch
	// a more experienced pokemon is harder to catch
	const threshold = 50                         // easy mode, typical base level is +/- 30
	randNum := rand.Intn(pokemon.BaseExperience) // generates a random number between 0 -> n
	fmt.Println("base experience:", pokemon.BaseExperience, "randNum:", randNum, "threshold:", threshold)
	if randNum > threshold {
		// pokemon is not caught
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon // inserts the pokemon into the pokemon map
	fmt.Printf("Succeeded in catching %s!\n", pokemonName)

	return nil
}
