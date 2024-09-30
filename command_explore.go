package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location provided")
	}
	area := args[0]
	resp, err := cfg.pokeapiClient.ListLocationAreasPokemon(&area)
	if err != nil {
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
