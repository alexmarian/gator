package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/state"
)

func HandleLogin(state *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("<username> is required")
	}
	user, err := state.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user does not exist: %v", err)
	}
	fmt.Printf("Running command %v", cmd.Args)
	err = state.Config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Printf("User set %s\n", cmd.Args[0])
	return nil
}
