package commands

import (
	"fmt"
)

func commandHelp(cfg *Config, _ []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}
