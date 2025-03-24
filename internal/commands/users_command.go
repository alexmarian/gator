package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/state"
)

func HandleUsers(state *state.State, cmd Command) error {

	users, err := state.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}
	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if state.Config.CurrentUserName == user.Name {
			fmt.Printf(" (current)")
		}
		fmt.Println("\n")
	}
	return nil
}
