package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCmds() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    cmdExit,
		},
		"help": {
			name:        "help",
			description: "displays an help message",
			callback:    cmdHelp,
		},
		"map": {
			name:        "map",
			description: "prints next 20 locations",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "prints previous 20 locations",
			callback:    cmdMapb,
		},
	}
}

func cmdExit(c *config) error {
	fmt.Println("exiting...")
	os.Exit(0)
	return nil
}

func cmdHelp(c *config) error {
	fmt.Println("welcome to godex!")
	fmt.Println("usage:")
	for _, cmd := range getCmds() {
		fmt.Printf("• %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func cmdMap(c *config) error {
	res, err := c.client.PrintLocationArea(c.nextMapUrl)
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

func cmdMapb(c *config) error {
	if c.prevMapUrl == nil {
		return errors.New("you're on the first page")
	}
	res, err := c.client.PrintLocationArea(c.prevMapUrl)
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
