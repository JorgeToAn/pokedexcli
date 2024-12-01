package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a list of available commands",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCommand,
		},
	}
}

func helpCommand() {
	availableCommands := getCommands()

	fmt.Println("Available commands:")
	for _, command := range availableCommands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
}

func exitCommand() {
	fmt.Println("Closing pokedex...")
	os.Exit(0)
}
