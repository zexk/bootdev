package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore <location>",
			description: "prints mon found in location",
			callback:    cmdExplore,
		},
		"catch": {
			name:        "catch <mon_name>",
			description: "attempts to capture mon",
			callback:    cmdCatch,
		},
		"inspect": {
			name:        "inspect <mon_name>",
			description: "inspects mon",
			callback:    cmdInspect,
		},
		"dex": {
			name:        "dex",
			description: "lists mon in dex",
			callback:    cmdDex,
		},
	}
}
