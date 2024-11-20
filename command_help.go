package main

import (
	"fmt"

	"github.com/st5/gator/commands"
)

func CallbackHelp(state commands.State, params ...string) error{

	cmds := getCommands()

	for _, cmd := range cmds {
		fmt.Printf("%s      %s\n" , cmd.name, cmd.description)	
		fmt.Println()
	}
	return nil
}