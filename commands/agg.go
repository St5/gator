package commands

import (
	"context"
	"fmt"

	"github.com/st5/gator/internal/rss"
)

func CallbackAgg(state State, params ...string) error{

	const url = "https://www.wagslane.dev/index.xml" 

	listPosts, err := rss.FetchFeed(context.Background(), url)

	if err != nil {
		return err
	}

	
	fmt.Printf("%v", listPosts)

	return nil
}