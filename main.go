package main

import (
	"time"

	"github.com/xristoskrik/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient           pokeapi.Client
	previousLocationAreaUrl *string
	nextLocationAreaUrl     *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cfg := Config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
