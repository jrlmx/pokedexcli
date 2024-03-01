package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	commands := getCommands()
	fmt.Println("Available commands:")
	for _, command := range commands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	return nil
}
