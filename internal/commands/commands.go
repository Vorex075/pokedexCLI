package commands

import (
	"fmt"
	"os"
)

var AllowedCommands map[string]CliCommand

func init() {
	AllowedCommands = map[string]CliCommand{
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
	Callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range AllowedCommands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func commandMap() error {
	return nil
}
