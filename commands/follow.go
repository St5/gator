package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/st5/gator/internal/database"
)

func CallbackFollow(state State, params ...string) error{

	if len(params) < 1 {
		return errors.New("url is required param")
	}

	userName := state.Config.CurrentUserName

	user, err := state.Db.GetUser(context.Background(), userName)

	if err != nil {
		return err
	}

	userId := user.ID

	url := params[0]

	feed, err := state.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	ffParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: userId,
		FeedID: feed.ID,
	}

	_, err = state.Db.CreateFeedFollow(context.Background(), ffParams)
	
	if err != nil {
		return err
	}

	fmt.Println("The feed was following")

	return nil

}