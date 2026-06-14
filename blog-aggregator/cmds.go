package main

import "errors"

type command struct {
	Name string
	Args []string
}

type cmds struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *cmds) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *cmds) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
