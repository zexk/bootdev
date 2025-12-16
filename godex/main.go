package main

import (
	"github.com/zexk/godex/internal/api"
	"time"
)

func main() {
	cfg := &config{
		client: api.NewClient(5*time.Second, 5*time.Second),
	}
	replLoop(cfg)
}
