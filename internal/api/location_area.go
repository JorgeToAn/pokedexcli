package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getNextURL(c *Config) string {
	if c.Next == nil {
		return "https://pokeapi.co/api/v2/location-area"
	}
	return *c.Next
}

func getPrevURL(c *Config) (string, error) {
	if c.Previous == nil {
		return "", fmt.Errorf("no previous areas")
	}
	return *c.Previous, nil
}

func GetNextLocationAreas(c *Config) ([]LocationArea, error) {
	nextURL := getNextURL(c)
	resp, err := http.Get(nextURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := locationAreasResponse{}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	c.Next = result.Next
	c.Previous = result.Previous

	return result.Results, nil
}

func GetPreviousLocationAreas(c *Config) ([]LocationArea, error) {
	prevURL, err := getPrevURL(c)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(prevURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := locationAreasResponse{}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	c.Next = result.Next
	c.Previous = result.Previous

	return result.Results, nil
}
