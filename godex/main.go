package main

import (
	"github.com/zexk/godex/internal/api"
	"time"
)

func main() {
	client := api.NewClient(5 * time.Second)
	cfg := &config{
		client: client,
	}
	replLoop(cfg)
}
