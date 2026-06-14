package main

import (
	"fmt"
)

func cmdHelp(c *config, args ...string) error {
	fmt.Println("welcome to godex!")
	fmt.Println("usage:")
	for _, cmd := range getCmds() {
		fmt.Printf("• %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
