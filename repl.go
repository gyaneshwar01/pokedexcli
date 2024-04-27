package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCliCommands()

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		cleanedText := cleanInput(text)

		if len(cleanedText) == 0 {
			continue
		}

		commandName := cleanedText[0]

		if command, ok := commands[commandName]; ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("invalid command")
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page location areas",
			callback:    callbackMapb,
		},
	}
}
