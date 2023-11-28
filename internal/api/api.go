package api

import (
	"net/http"
	"time"

	"github.com/ClemSK/repl/internal/cache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
