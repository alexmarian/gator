package commands

import (
	"fmt"
	"github.com/alexmarian/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Commands map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	if c.Commands == nil {
		c.Commands = make(map[string]func(*state.State, Command) error)
	}
	c.Commands[name] = f
}
func (c *Commands) Run(s *state.State, cmd Command) error {
	if f, ok := c.Commands[cmd.Name]; ok {
		return f(s, cmd)
	} else {
		return fmt.Errorf("command %s not found", cmd.Name)
	}
}
