package main

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("No caught pokemon yet")
		return nil
	}
	for _, v := range cfg.caughtPokemon {
		fmt.Println(v.Name)
	}
	return nil
}
