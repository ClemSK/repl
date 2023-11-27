package main

import "github.com/ClemSK/repl/internal/api"

type config struct {
	pokeApiClient       api.Client
	nextLocationAreaUrl *string
	// prevLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeApiClient: api.NewClient(),
	}
	startRepl(&cfg)
}
