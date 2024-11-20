package commands

import (
	"context"
	"strconv"

	"github.com/st5/gator/internal/database"
)

func CallbackBrowse(state State, user database.User, params ...string) error {

	limit := int32(2)
	if len(params) > 0 {
		if limitP, err := strconv.ParseInt(params[0],10,10); err == nil {
			limit = int32(limitP)
		}
	}

	param := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := state.Db.GetPostsForUser(context.Background(), param)
	if err != nil {
		return err
	}

	for _, post := range posts {
		println(post.Title)
	}
	return nil
}
