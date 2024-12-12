package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JorgeToAn/pokedexcli/internal/api"
)

type Config struct {
	Next      *string
	Previous  *string
	ApiClient api.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, args []string) error
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	// REPL loop
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		name, args := cleanInput(scanner.Text())
		command, exists := commands[name]
		if !exists {
			fmt.Printf("Command '%s' not found, use 'help' to see a list of available commands\n", name)
			continue
		}

		err := command.callback(config, args)
		if err != nil {
			fmt.Println(err)
		}
	}
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
		"explore": {
			name:        "explore",
			description: "Shows the Pokemon in the given area",
			callback:    exploreCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCommand,
		},
	}
}

func cleanInput(input string) (string, []string) {
	lowered := strings.ToLower(input)
	trimmed := strings.Trim(lowered, " ")
	parts := strings.Split(trimmed, " ")

	name := parts[0]
	args := parts[1:]
	return name, args
}
