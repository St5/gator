package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/st5/gator/internal/database"
)

func CallbackRegister(state State, params ...string) error{

	if len(params) == 0 {
		return errors.New("a username is required")
	}
	username := params[0]

	_, err := state.Db.GetUser(context.Background(), username)
	
	if err == nil {
		return errors.New("User is present already")
	}


	userParam := database.CreateUserParams{
		ID : uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,    
	}
	_ , err = state.Db.CreateUser(context.Background(), userParam)

	if err != nil {
		return err
	}
	fmt.Println("User was created")
	CallbackLogin(state, username)

	

	return nil
}