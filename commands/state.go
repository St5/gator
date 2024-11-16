package commands

import (
	"github.com/st5/gator/internal/config"
	"github.com/st5/gator/internal/database"
)


type State struct {
	Config config.Config
	Db *database.Queries
}