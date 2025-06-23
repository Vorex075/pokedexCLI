package commands

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapb,
		},
	}
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
}
