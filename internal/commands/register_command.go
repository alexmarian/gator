package commands

import (
	"context"
	"fmt"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
	"github.com/google/uuid"
	"time"
)

func HandleRegister(state *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("<username> is required")
	}
	fmt.Printf("Running command %v", cmd.Args)
	_, err := state.Db.GetUser(context.Background(), cmd.Args[0])
	if err == nil {
		return fmt.Errorf("error user already exists")
	}

	cup := database.CreateUserParams{ID: uuid.New(), Name: cmd.Args[0], CreatedAt: time.Now(), UpdatedAt: time.Now()}
	createUser, err := state.Db.CreateUser(context.Background(), cup)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	state.Config.SetUser(createUser.Name)
	fmt.Printf("User registered %v\n", createUser)
	return nil
}
