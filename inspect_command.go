package main

import (
	"errors"
	"fmt"
)

func inspectCommand(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon>")
	}

	pokemon, exists := config.CaughtPokemon[args[0]]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}

	return nil
}
