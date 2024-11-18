package rss

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error){

	request, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)

	if err !=nil {
		return nil, err
	}

	request.Header.Add("User-Agent", "gator")

	client := http.Client{
		Timeout: time.Minute,
	}

	res, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err!=nil {
		return nil, err
	}

	rssFeed := RSSFeed {}

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil


}

