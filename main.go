package main

import (
	"bufio"
	"fmt"
	"github.com/Vorex075/pokedexCLI/internal/commands"
	"github.com/Vorex075/pokedexCLI/internal/parsing"
	"log"
	"os"
	"strings"
)

func main() {
	lineReader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !lineReader.Scan() {
			log.Fatal("error in scan")
		}
		allText := lineReader.Text()
		separatedWords := parsing.CleanInput(allText)
		if len(separatedWords) == 0 {
			continue
		}
		command, ok := commands.AllowedCommands[separatedWords[0]]
		if ok {
			if err := command.Callback(); err != nil {
				fmt.Printf("error while trying to execute %s command: %v",
					separatedWords[0], err)
				return
			}
		} else {
			fmt.Printf("Your command was: %s\n", strings.ToLower(separatedWords[0]))
		}
	}
}
