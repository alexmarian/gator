package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
)

func HandleFeeds(state *state.State, cmd Command) error {
	feeds, err := state.Db.GetAllFeedsWithUserNames(context.Background())
	if err != nil {
		return fmt.Errorf("error creating feeds: %v", err)
	}

	fmt.Println("Feeds fetched successfully:")
	for _, feed := range feeds {
		printFeedExcerpt(feed)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeedExcerpt(feed database.GetAllFeedsWithUserNamesRow) {
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* Username:        %s\n", feed.UserName)
}
