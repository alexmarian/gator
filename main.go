package main

import (
	"github.com/alexmarian/gator/internal/commands"
	"github.com/alexmarian/gator/internal/config"
	"github.com/alexmarian/gator/internal/state"
	"log"
	"os"
)

const configUserName = "lane"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	state := state.State{
		Config: cfg,
	}
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <command> [args]", os.Args[0])
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := commands.Command{
		Name: cmdName,
		Args: cmdArgs,
	}
	err = getCommands().Run(&state, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}
}

func getCommands() *commands.Commands {
	cmds := commands.Commands{}
	cmds.Register("login", commands.HandleLogin)
	return &cmds
}
