package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error making request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error getting response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return &RSSFeed{}, fmt.Errorf("unexpected status: %s\n", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading body: %w", err)
	}
	defer res.Body.Close()

	feed := &RSSFeed{}
	err = xml.Unmarshal(body, feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshaling xml: %w", err)
	}

	return feed, nil
}

func unescapeHtml(f *RSSFeed) {
	html.UnescapeString(f.Channel.Title)
	html.UnescapeString(f.Channel.Description)

	for _, item := range f.Channel.Item {
		html.UnescapeString(item.Title)
		html.UnescapeString(item.Description)
	}
}
