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
	action      func(*config) error
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

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.action(cfg)

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
	}
}
