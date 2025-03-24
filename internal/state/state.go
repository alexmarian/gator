package state

import (
	"github.com/alexmarian/gator/internal/config"
	"github.com/alexmarian/gator/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}
