package main

import (
	"time"

	"github.com/JorgeToAn/pokedexcli/internal/api"
)

func main() {
	config := Config{
		ApiClient: api.NewClient(time.Second * 30),
	}

	startRepl(&config)
}
