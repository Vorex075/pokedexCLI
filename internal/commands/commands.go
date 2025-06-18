package commands

func getCommands() map[string]CliCommand {
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
	}
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
}
