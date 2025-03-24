package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	requestWithContext, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	requestWithContext.Header.Set("User-Agent", "gator")
	httpClient := &http.Client{}
	response, err := httpClient.Do(requestWithContext)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return parseFeed(response.Body)
}

func parseFeed(body io.ReadCloser) (*RSSFeed, error) {
	all, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	rssFeed := new(RSSFeed)
	err = xml.Unmarshal(all, rssFeed)
	if err != nil {
		return nil, err
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}
	return rssFeed, nil
}
