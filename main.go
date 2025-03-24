package main

import (
	"database/sql"
	"github.com/alexmarian/gator/internal/commands"
	"github.com/alexmarian/gator/internal/config"
	"github.com/alexmarian/gator/internal/database"
	"github.com/alexmarian/gator/internal/state"
	"log"
	"os"
)
import _ "github.com/lib/pq"

const configUserName = "lane"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	dbQueries := database.New(db)
	state := state.State{
		Config: cfg,
		Db:     dbQueries,
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
	cmds.Register("register", commands.HandleRegister)
	cmds.Register("reset", commands.HandleReset)
	cmds.Register("users", commands.HandleUsers)
	cmds.Register("agg", commands.HandleAgg)
	cmds.Register("addfeed", commands.HandleAddFeed)
	cmds.Register("feeds", commands.HandleFeeds)
	return &cmds
}
