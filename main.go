package main

import (
	"github.com/kachit/golang-api-skeleton/commands/application"
	"github.com/kachit/golang-api-skeleton/commands/database"
	"github.com/kachit/golang-api-skeleton/commands/develop"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "golang-api-skeleton",
		Commands: []*cli.Command{
			//database
			database.NewDatabaseMigrateCommand(),
			database.NewDatabaseRollbackCommand(),
			database.NewDatabaseSeedCommand(),
			database.NewDatabaseClearCommand(),
			//develop
			develop.NewDevelopTestCommand(),
			//application
			application.NewApplicationStartCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
