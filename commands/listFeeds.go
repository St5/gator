package commands

import (
	"context"
	"fmt"
)

func CallbackListFeeds(state State, params ...string) error{

	feeds, err := state.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s - %s \n", feed.Name, feed.Name_2.String)
	}

 return nil
}