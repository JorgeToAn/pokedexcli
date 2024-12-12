package main

import (
	"errors"
	"fmt"
)

func exploreCommand(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <area-name>")
	}
	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	locationAreaDetail, err := config.ApiClient.GetLocationAreaDetail(areaName)
	if err != nil {
		return err
	}

	for _, encounter := range locationAreaDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
