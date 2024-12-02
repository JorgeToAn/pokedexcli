package main

import (
	"fmt"
	"os"

	"github.com/JorgeToAn/pokedexcli/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *api.Config)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a list of available commands",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 areas",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas",
			callback:    mapbCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCommand,
		},
	}
}

func helpCommand(c *api.Config) {
	availableCommands := getCommands()

	fmt.Println("Available commands:")
	for _, command := range availableCommands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
}

func mapCommand(c *api.Config) {
	locations, err := api.GetNextLocationAreas(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, area := range locations {
		fmt.Printf("  %s\n", area.Name)
	}
}

func mapbCommand(c *api.Config) {
	locations, err := api.GetPreviousLocationAreas(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, area := range locations {
		fmt.Printf("  %s\n", area.Name)
	}
}

func exitCommand(c *api.Config) {
	fmt.Println("Closing pokedex...")
	os.Exit(0)
}
