package pokeapi

import (
	"net/http"
	"time"

	"github.com/xristoskrik/pokedexcli/internal/pokecache"
)

const BaseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.PokeCache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
