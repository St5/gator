package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/st5/gator/internal/database"
	"github.com/st5/gator/internal/rss"
)

func CallbackAgg(state State, params ...string) error{

	// const url = "https://www.wagslane.dev/index.xml" 

	if len(params) < 1 {
		return fmt.Errorf("duration is required param")
	}
	duration, err := time.ParseDuration(params[0])

	if err != nil {
		return err
	}

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err := scrapeFeeds(state)
		if err != nil {
			return err
		}
	}

	return nil
}


func scrapeFeeds(state State) error {

	feed, err := state.Db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return err
	}

	listPosts, err := rss.FetchFeed(context.Background(), feed.Url)

	if err != nil {
		return err
	}

	i := 0
	fmt.Printf("============= %s =============\n", listPosts.Channel.Title)
	for _, post := range listPosts.Channel.Item {

		published, _ := time.Parse(time.RFC1123, post.PubDate)
		postParam := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       post.Title,
			Url:         post.Link,
			Description: sql.NullString{String: post.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: published, Valid: true},
			FeedID:      feed.ID,
		}
		_, err := state.Db.CreatePost(context.Background(), postParam)

		if err == nil {
			i++
		}
	}
	fmt.Printf("%d new posts\n", i)
	fmt.Println("==================================")	

	err = state.Db.MarkFeedFetched(context.Background(), feed.ID)

	if err != nil {
		return err
	}
	
	return nil
}