package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
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
		"map": {
			name:        "map",
			description: "goes 20 location forward",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "goes 20 location backward",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "explore  {name or id}",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch {name or id}",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect pokemon",
			description: "inspect {name or id}",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
