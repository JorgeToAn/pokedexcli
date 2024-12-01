package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

		command.callback()
	}
}
