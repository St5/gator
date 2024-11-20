package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/st5/gator/internal/database"
)

func CallbackUnfollow(state State, user database.User, params ...string) error {

	if len(params) < 1 {
		return errors.New("url is required param")
	}
	userId := user.ID

	url := params[0]

	feed, err := state.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	paramsWhere := database.UnfollowParams {
		UserID: userId,
		FeedID: feed.ID,
	}

	err = state.Db.Unfollow(context.Background(), paramsWhere)

	if err != nil {
		return err
	}

	fmt.Println("The feed was unfollowing")

	return nil
}