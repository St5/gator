package commands

import (
	"context"
	"fmt"
)

func CallbackUsers(state State, params ...string) error{

	users, err := state.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := state.Config.CurrentUserName

	for _, user := range users {

		add:=""
		if user.Name == currentUser {
			add = "(current)"
		}
		fmt.Printf("%s %s\n", user.Name, add)
	}
	return nil
}