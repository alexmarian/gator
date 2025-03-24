package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/rss"
	"github.com/alexmarian/gator/internal/state"
)

const feedUrl = "https://www.wagslane.dev/index.xml"

func HandleAgg(state *state.State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}
	fmt.Printf("Feed fetched: %v\n", feed)
	return nil
}
