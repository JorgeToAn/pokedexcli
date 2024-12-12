package main

import (
	"fmt"
	"os"
)

func exitCommand(config *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
