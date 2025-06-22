package commands

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "Help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "Map",
			Description: "Get the next page of locations",
			Callback:    commandMap,
		},
	}
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
}
