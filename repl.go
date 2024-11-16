package main

import (
	"fmt"
	"os"
	"github.com/st5/gator/commands"
)

type cliCommands struct {
	name string
	description string
	callback func (state commands.State, params ...string) error
}

func startRepl(state commands.State){

	avilableComands := getCommands()

	args := os.Args

	if len(args) < 2 {
		fmt.Println(" not enough arguments were provided")
		os.Exit(1)
	}

	commandName := args[1]

	command, ok := avilableComands[commandName]

	if !ok {
		fmt.Println("invalid command")
		os.Exit(1)
	}

	params := args[2:]

	err := command.callback(state, params...)

	if err != nil {
		fmt.Printf("Error %v \n", err)
		os.Exit(1)
	}

}

func getCommands() map[string]cliCommands {

	return map[string]cliCommands {
		"help": {
			name: "help",
			description: "Help commands",
			callback: commands.CallbackHelp,
		},
		"login": {
			name:        "login",
			description: "Login as user",
			callback:    commands.CallbackLogin,
		},
		"register": {
			name:        "register",
			description: "Register an user",
			callback:    commands.CallbackRegister,
		},
		"reset": {
			name:        "rest",
			description: "Clear user table",
			callback:    commands.CallbackReset,
		},
		"users": {
			name:        "users",
			description: "Return list of users",
			callback:    commands.CallbackUsers,
		},
	}
}