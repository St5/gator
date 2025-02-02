package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/st5/gator/internal/database"
)

func CallbackAddFeed(state State, user database.User, params ...string) error{

	if len(params) < 2 {
		return errors.New("feed name and url is required param")
	}

	userId := user.ID

	name := params[0]
	url := params[1]

	feedParam := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: userId,
	}

	_ , err := state.Db.CreateFeed(context.Background(), feedParam)

	if err != nil {
		return err
	}

	fmt.Println("Feed was saved")

	err = CallbackFollow(state, user, url)

	if err != nil {
		return err
	}

	return nil
}