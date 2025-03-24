package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
)

func MiddlewareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
	return func(state *state.State, cmd Command) error {
		user, err := state.Db.GetUser(context.Background(), state.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		}
		return handler(state, cmd, user)
	}
}
