package main

import (
	"fmt"
)

func helpCommand(c *Config, args []string) error {
	availableCommands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range availableCommands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	return nil
}
