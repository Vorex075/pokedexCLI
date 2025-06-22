package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Vorex075/pokedexCLI/internal/api"
	"github.com/Vorex075/pokedexCLI/internal/commands"
	"github.com/Vorex075/pokedexCLI/internal/parsing"
)

func main() {
	lineReader := bufio.NewScanner(os.Stdin)
	cfg := commands.NewConfig(api.NewClient(5 * time.Second))
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
		allowedCommands := commands.GetCommands()
		command, ok := allowedCommands[separatedWords[0]]
		if ok {
			if err := command.Callback(&cfg); err != nil {
				fmt.Printf("error while trying to execute %s command: %v",
					separatedWords[0], err)
				return
			}
		} else {
			fmt.Printf("Your command was: %s\n", strings.ToLower(separatedWords[0]))
		}
	}
}
