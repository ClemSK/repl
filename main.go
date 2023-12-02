package main

import (
	"time"

	"github.com/ClemSK/repl/internal/api"
)

type config struct {
	pokeApiClient       api.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	caughtPokemon       map[string]api.Pokemon
}

func main() {
	cfg := config{
		pokeApiClient: api.NewClient(time.Hour),
		caughtPokemon: make(map[string]api.Pokemon),
	}
	startRepl(&cfg)
}
