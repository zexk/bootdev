package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cmds = map[string]cliCommand{
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
}

func cmdExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cmdHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	//TODO parse cmd usage from cmds
	return nil
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	words := strings.Split(text, " ")
	return words
}

func replLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		fmt.Print("godex ~ ")
		scanner.Scan()
		input += scanner.Text()
		words := cleanInput(input)
		if cmd, found := cmds[words[0]]; found {
			cmd.callback()
		} else {
			fmt.Printf("%s is not a valid command\n", words[0])
		}
		input = ""
	}
}
