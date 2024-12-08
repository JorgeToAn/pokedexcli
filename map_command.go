package main

import (
	"errors"
	"fmt"
)

func mapCommand(c *Config, args []string) error {
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

func mapbCommand(c *Config, args []string) error {
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
