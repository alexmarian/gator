package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
)

func HandleFollowing(state *state.State, cmd Command, user database.User) error {
	feeds, err := state.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error creating feeds: %v", err)
	}

	fmt.Printf("Feeds for %s fetched successfully:\n", user.Name)
	for _, feed := range feeds {
		printFeedFollowExcerpt(feed)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeedFollowExcerpt(feed database.GetFeedFollowsForUserRow) {
	fmt.Printf("* Name:          %s\n", feed.FeedName)
}
