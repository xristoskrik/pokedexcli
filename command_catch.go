package main

import (
	"errors"
	"fmt"

	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.PokemonInfo(&pokemonName)
	if err != nil {
		return err
	}
	const thresHold = 80
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > thresHold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("caught %s\n", pokemonName)

	return nil
}
