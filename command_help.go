package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, command := range getCliCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()

	return nil
}
