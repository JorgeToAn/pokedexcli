package api

import (
	"net/http"
	"time"

	"github.com/JorgeToAn/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: *http.DefaultClient,
		cache:      pokecache.NewCache(interval),
	}
}
