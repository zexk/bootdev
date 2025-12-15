package api

import (
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type Client struct {
	client http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
	}
}
