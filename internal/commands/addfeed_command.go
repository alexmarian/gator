package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
	"github.com/google/uuid"
	"time"
)

func HandleAddFeed(state *state.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("<name> <url> are required")
	}
	user, err := state.Db.GetUser(context.Background(), state.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user does not exist: %v", err)
	}
	feedParam := database.CreateFeedParams{ID: uuid.New(), Name: cmd.Args[0], CreatedAt: time.Now(), UpdatedAt: time.Now(), Url: cmd.Args[1], UserID: user.ID}
	feed, err := state.Db.CreateFeed(context.Background(), feedParam)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
