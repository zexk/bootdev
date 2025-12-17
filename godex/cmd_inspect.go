package main

import (
	"fmt"
)

func cmdInspect(c *config, args ...string) error {
	if mon, found := c.dex[args[0]]; found {
		fmt.Printf("name: %s\n", mon.Name)
		fmt.Printf("height: %d\n", mon.Height)
		fmt.Printf("weight: %d\n", mon.Weight)
		fmt.Println("stats:")
		for _, stat := range mon.Stats {
			fmt.Printf("• %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("types:")
		for _, typ := range mon.Types {
			fmt.Printf("• %s\n", typ.Type.Name)
		}
	} else {
		fmt.Println(" you have not caught that mon")
	}
	return nil
}
