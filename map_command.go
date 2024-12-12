package main

import (
	"errors"
	"fmt"
)

func mapCommand(config *Config, args ...string) error {
	locationAreasResp, err := config.ApiClient.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationAreasResp.Next
	config.Previous = locationAreasResp.Previous

	for _, area := range locationAreasResp.Results {
		fmt.Printf("  %s\n", area.Name)
	}
	return nil
}

func mapbCommand(config *Config, args ...string) error {
	if config.Previous == nil {
		return errors.New("no previous areas")
	}

	locationAreasResp, err := config.ApiClient.GetLocationAreas(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationAreasResp.Next
	config.Previous = locationAreasResp.Previous

	for _, area := range locationAreasResp.Results {
		fmt.Printf("  %s\n", area.Name)
	}
	return nil
}
