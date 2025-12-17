package main

import (
	"github.com/zexk/godex/internal/api"
	"time"
)

func main() {
	cfg := &config{
		dex : map[string]api.MonRes{},
		client: api.NewClient(5*time.Second, 5*time.Minute),
	}
	replLoop(cfg)
}
