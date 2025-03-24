package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
	"github.com/google/uuid"
	"time"
)

func HandleFollow(state *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("<url> is required")
	}
	feed, err := state.Db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("feed does not exist: %v", err)
	}
	feedFollowParams := database.CreateFeedFollowParams{ID: uuid.New(), UserID: user.ID, FeedID: feed.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	feedFollow, err := state.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("* ID:            %s\n", feedFollow.ID)
	fmt.Printf("* Created:       %v\n", feedFollow.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feedFollow.UpdatedAt)
	fmt.Printf("* Feed Name:          %s\n", feedFollow.FeedName)
	fmt.Printf("* User Name:           %s\n", feedFollow.UserName)
}
