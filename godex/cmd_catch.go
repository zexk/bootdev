package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func cmdCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a name")
	}
	res, err := c.client.GetMon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("throwing a ball at %s...\n", args[0])
	if rand.Intn(res.BaseExperience) > 40 {
		fmt.Printf("%s escaped!\n", args[0])
		return nil
	} else {
		fmt.Printf("%s was caught!\n", args[0])
		c.dex[args[0]] = res
	}
	return nil
}
