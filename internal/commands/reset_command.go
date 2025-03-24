package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/state"
)

func HandleReset(state *state.State, cmd Command) error {
	err := state.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting all users: %v", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
