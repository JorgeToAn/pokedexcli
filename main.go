package main

import (
	"time"

	"github.com/JorgeToAn/pokedexcli/internal/api"
)

func main() {
	config := Config{
		ApiClient:     api.NewClient(time.Second * 30),
		CaughtPokemon: make(map[string]api.PokemonResponse),
	}

	startRepl(&config)
}
