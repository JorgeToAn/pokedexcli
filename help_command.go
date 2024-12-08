package main

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, args []string) error
}

func helpCommand(c *Config, args []string) error {
	availableCommands := getCommands()

	fmt.Println("Available commands:")
	for _, command := range availableCommands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	return nil
}
