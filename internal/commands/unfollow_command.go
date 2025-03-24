package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
)

func HandleUnfollow(state *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("<url> is required")
	}
	feed, err := state.Db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("feed does not exist: %v", err)
	}
	deleteFeedFollowParams := database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID}
	err = state.Db.DeleteFeedFollow(context.Background(), deleteFeedFollowParams)
	if err != nil {
		return fmt.Errorf("error deleting feed follow: %v", err)
	}

	fmt.Println("Feed unfollowed successfully:")
	return nil
}
