package commands

import (
	"context"
	"errors"
	"fmt"
)

func CallbackLogin(state State, params ...string) error{

	if len(params) == 0 {
		return errors.New("a username is required")
	}

	username := params[0]

	_, err := state.Db.GetUser(context.Background(), username)
	
	if err != nil {
		return errors.New("User absend in db")
	}

	err = state.Config.SetUser(username)

	if err != nil {
		return err
	}

	fmt.Println("User was login")
	return nil
}