package main

import (
	"errors"
	"fmt"
)

func cmdMap(c *config, args ...string) error {
	res, err := c.client.GetLocationArea(c.nextMapUrl)
	if err != nil {
		return err
	}

	c.nextMapUrl = res.Next
	c.prevMapUrl = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func cmdMapb(c *config, args ...string) error {
	if c.prevMapUrl == nil {
		return errors.New("you're on the first page")
	}
	res, err := c.client.GetLocationArea(c.prevMapUrl)
	if err != nil {
		return err
	}

	c.nextMapUrl = res.Next
	c.prevMapUrl = res.Previous

	for _, location := range res.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
