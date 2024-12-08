package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JorgeToAn/pokedexcli/internal/api"
)

type Config struct {
	Next      *string
	Previous  *string
	ApiClient api.Client
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	// REPL loop
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		command, exists := commands[input]
		if !exists {
			fmt.Printf("Command '%s' not found, use 'help' to see a list of available commands\n", input)
			continue
		}

		err := command.callback(config)
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
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCommand,
		},
	}
}
