package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

func helpCommand(c *Config) error {
	availableCommands := getCommands()

	fmt.Println("Available commands:")
	for _, command := range availableCommands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	return nil
}

func mapCommand(c *Config) error {
	locationAreasResp, err := c.ApiClient.GetLocationAreas(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationAreasResp.Next
	c.Previous = locationAreasResp.Previous

	for _, area := range locationAreasResp.Results {
		fmt.Printf("  %s\n", area.Name)
	}
	return nil
}

func mapbCommand(c *Config) error {
	if c.Previous == nil {
		return errors.New("no previous areas")
	}

	locationAreasResp, err := c.ApiClient.GetLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locationAreasResp.Next
	c.Previous = locationAreasResp.Previous

	for _, area := range locationAreasResp.Results {
		fmt.Printf("  %s\n", area.Name)
	}
	return nil
}

func exitCommand(c *Config) error {
	fmt.Println("Closing pokedex...")
	os.Exit(0)
	return nil
}
