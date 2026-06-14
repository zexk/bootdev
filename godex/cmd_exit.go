package main

import (
	"fmt"
	"os"
)

func cmdExit(c *config, args ...string) error {
	fmt.Println("exiting...")
	os.Exit(0)
	return nil
}
