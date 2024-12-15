package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchCommand(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := config.ApiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	catchRand := rand.Float32()
	// 1 - (BaseExperience / 500) = SuccessRate
	if catchRand > float32(pokemon.BaseExperience)/500 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.CaughtPokemon[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
