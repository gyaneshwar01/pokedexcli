package main

import "github.com/gyaneshwar01/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiclient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := &config{
		pokeapiclient: pokeapi.NewClient(),
	}
	startRepl(cfg)
}
