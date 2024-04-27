package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCliCommands()

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		if command, ok := commands[text]; ok {
			command.callback()
		} else {
			fmt.Println("Invalid command, try again!")
		}
	}
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, command := range getCliCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
