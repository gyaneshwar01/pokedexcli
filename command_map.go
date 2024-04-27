package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	res, err := cfg.pokeapiclient.ListLocationArea(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	res, err := cfg.pokeapiclient.ListLocationArea(cfg.prevLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas: ")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}
