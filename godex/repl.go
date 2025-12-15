package main

import (
	"bufio"
	"fmt"
	"github.com/zexk/godex/internal/api"
	"os"
	"strings"
)

type config struct {
	client     api.Client
	nextMapUrl *string
	prevMapUrl *string
}

func replLoop(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		fmt.Print("godex ~ ")

		scanner.Scan()
		input += scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		if cmd, found := getCmds()[words[0]]; found {
			err := cmd.callback(c)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("%s is not a valid command\n", words[0])
		}

		input = ""
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
