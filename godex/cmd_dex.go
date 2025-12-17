package main

import (
	"fmt"
)

func cmdDex(c *config, args ...string) error {
	fmt.Println("your dex:")
	for _, entry := range c.dex {
		fmt.Printf("• %s\n", entry.Name)
	}
	return nil
}
