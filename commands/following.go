package commands

import (
	"context"
	"fmt"

	"github.com/st5/gator/internal/database"
)

func CallbackFollowing(state State, user database.User, params ...string) error{

	userName := state.Config.CurrentUserName

	user, err := state.Db.GetUser(context.Background(), userName)

	if err != nil {
		return err
	}

	userId := user.ID

	feeds, err := state.Db.GetFeedFollowsForUser(context.Background(), userId)

	if err != nil {
		return err
	}

	for _, feed :=range feeds {

		fmt.Println(feed.Name)
	}

	return nil


}