package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	action      func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")

		scanner.Scan()
		input := scanner.Text()

		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		commands := getCommands()

		command, ok := commands[cleaned[0]]

		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.action(cfg, args...)

		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func cleanInput(input string) []string {
	lowercase := strings.ToLower(input)
	words := strings.Fields(lowercase)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			action:      commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			action:      commandExit,
		},
		"map": {
			name:        "map",
			description: "List next page of location areas",
			action:      commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous page of location areas",
			action:      commandMapBack,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List pokemon in a location area",
			action:      commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attept to catch a pokemon and add it to your pokedex",
			action:      commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspect a pokemon in your pokedex",
			action:      commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemon in your pokedex",
			action:      commandPokedex,
		},
	}
}
