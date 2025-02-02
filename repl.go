package main

import (
	"context"
	"fmt"
	"os"

	"github.com/st5/gator/commands"
	"github.com/st5/gator/internal/database"
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
			callback: CallbackHelp,
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
		"agg": {
			name:        "agg <duration>",
			description: "Aggregate RSS",
			callback:    commands.CallbackAgg,
		},
		"addfeed": {
			name:        "addfeed <name> <url>",
			description: "Add feed to DB",
			callback:    middlewareLoggedIn(commands.CallbackAddFeed),
		},
		"feeds": {
			name:        "feeds",
			description: "List Feeds",
			callback:    commands.CallbackListFeeds,
		},
		"follow": {
			name:        "follow <url>",
			description: "Assign a user to a feed",
			callback:    middlewareLoggedIn(commands.CallbackFollow),
		},
		"following": {
			name:        "following",
			description: "Print all the names of the feeds the current user is following.",
			callback:    middlewareLoggedIn(commands.CallbackFollowing),
		},
		"unfollow": {
			name:        "unfollow <url>",
			description: "Print all the names of the feeds the current user is following.",
			callback:    middlewareLoggedIn(commands.CallbackUnfollow),
		},
		"browse": {
			name:        "browse <limit>",
			description: "Print all post of the feeds the current user is following.",
			callback:    middlewareLoggedIn(commands.CallbackBrowse),
		},
	}
}

func middlewareLoggedIn(handler func(state commands.State, user database.User, params ...string) error) func(state commands.State, params ...string) error {
	return func(state commands.State, params ...string) error {
		user, err := state.Db.GetUser(context.Background(), state.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(state, user, params...)
	}
}