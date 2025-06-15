package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	lineReader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !lineReader.Scan() {
			log.Fatal("error in scan")
		}
		allText := lineReader.Text()
		separatedWords := cleanInput(allText)
		if len(separatedWords) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", strings.ToLower(separatedWords[0]))
	}
}
