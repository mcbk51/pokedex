package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, i := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", i.Name)
	}
	return nil
}
