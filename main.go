package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/st5/gator/commands"
	"github.com/st5/gator/internal/config"
	"github.com/st5/gator/internal/database"
)



func main() {
	conf,err := config.Read()
	if err != nil {
		fmt.Printf("Error %v \n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", conf.DBURL)
	if err!= nil {
		fmt.Printf("Error connection: %v \n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	statePoint := commands.State{
		Config: conf,
		Db: dbQueries,
	}
	startRepl(statePoint)
}
