package main

import (
	"errors"
	"fmt"
)

func pokedexCommand(config *Config, args ...string) error {
	fmt.Println("Your Pokedex:")

	if len(config.CaughtPokemon) == 0 {
		return errors.New("you haven't caught any Pokemon")
	}

	for _, pokemon := range config.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
