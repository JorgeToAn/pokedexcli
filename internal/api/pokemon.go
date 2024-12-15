package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	value, exists := c.cache.Get(fullURL)
	if exists {
		result := PokemonResponse{}
		err := json.Unmarshal(value, &result)
		if err != nil {
			return PokemonResponse{}, err
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemon := PokemonResponse{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
