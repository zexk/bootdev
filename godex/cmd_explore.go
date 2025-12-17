package main

import (
	"errors"
	"fmt"
)

func cmdExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location area")

	}
	res, err := c.client.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Println("found mons:")
	for _, enc := range res.PokemonEncounters {
		fmt.Printf("• %s\n", enc.Pokemon.Name)
	}

	return nil
}
