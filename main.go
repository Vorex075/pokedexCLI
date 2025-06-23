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
	"github.com/Vorex075/pokedexCLI/internal/pokecache"
)

func main() {
	lineReader := bufio.NewScanner(os.Stdin)
	cfg := commands.NewConfig(api.NewClient(5*time.Second), pokecache.NewCache(5*time.Second))
	for {
		fmt.Print("Pokedex > ")
		if !lineReader.Scan() {
			log.Fatal("error in scan")
		}
		allText := lineReader.Text()
		commandAndArgs := parsing.CleanInput(allText)
		if len(commandAndArgs) == 0 {
			continue
		}
		allowedCommands := commands.GetCommands()
		command, ok := allowedCommands[commandAndArgs[0]]
		args := commandAndArgs[1:]
		if ok {
			if err := command.Callback(&cfg, args); err != nil {
				fmt.Printf("error while trying to execute %s command: %v",
					commandAndArgs[0], err)
				continue
			}
		} else {
			fmt.Printf("Your command was: %s\n", strings.ToLower(commandAndArgs[0]))
		}
	}
}
