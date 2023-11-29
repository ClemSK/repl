package api

import (
	"net/http"
	"time"

	"github.com/ClemSK/repl/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheTimeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheTimeout),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
