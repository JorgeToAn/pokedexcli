package api

import (
	"time"

	"github.com/JorgeToAn/pokedexcli/internal/pokecache"
)

type Config struct {
	Next     *string
	Previous *string
	cache    pokecache.Cache
}

func NewConfig() Config {
	return Config{
		Next:     nil,
		Previous: nil,
		cache:    pokecache.NewCache(time.Second * 30),
	}
}
