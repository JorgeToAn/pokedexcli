package main

import (
	"fmt"
	"math/rand"
)

const (
	baseCatchRate = 500
)

func catchCommand(config *Config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("usage: catch <pokemon>")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := config.ApiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	catchRand := rand.Intn(baseCatchRate)
	if catchRand > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
