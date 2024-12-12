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
	callback    func(config *Config, args ...string) error
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	// REPL loop
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		commandName := words[0]
		args := words[1:]

		command, exists := commands[commandName]
		if !exists {
			fmt.Println("unknown command")
			continue
		}

		err := command.callback(config, args...)
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

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	trimmed := strings.Trim(lowered, " ")
	words := strings.Split(trimmed, " ")

	return words
}
