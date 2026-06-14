package api

import (
	"github.com/zexk/godex/internal/cache"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type Client struct {
	cache  cache.Cache
	client http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *cache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: timeout,
		},
	}
}
