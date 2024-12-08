package main

import (
	"fmt"
	"os"
)

func exitCommand(c *Config, args []string) error {
	fmt.Println("Closing pokedex...")
	os.Exit(0)
	return nil
}
