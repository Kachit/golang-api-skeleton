package commands

import (
	"github.com/kachit/golang-api-skeleton/commands/database"
	"github.com/kachit/golang-api-skeleton/commands/develop"
	"github.com/kachit/golang-api-skeleton/commands/server"
	"github.com/mitchellh/cli"
)

type Factory func(cli.Ui) (cli.Command, error)

var registry map[string]Factory

func GetCmds(ui cli.Ui) map[string]cli.CommandFactory {
	res := make(map[string]cli.CommandFactory)
	for name, fn := range registry {
		fnc := fn
		res[name] = func() (cli.Command, error) {
			return fnc(ui)
		}
	}
	return res
}

func AddRegistry(name string, fn Factory) {
	if registry == nil {
		registry = make(map[string]Factory)
	}
	registry[name] = fn
}

func init() {
	//Develop
	AddRegistry(commands_develop.DevelopTest, commands_develop.DevelopTestCmd)
	//Database
	AddRegistry(commands_database.DatabaseMigrationsMigrate, commands_database.MigrateCmd)
	AddRegistry(commands_database.DatabaseMigrationsRollback, commands_database.RollbackCmd)
	AddRegistry(commands_database.DatabaseSeedersSeed, commands_database.SeedCmd)
	AddRegistry(commands_database.DatabaseSeedersClear, commands_database.ClearCmd)
	//Server
	AddRegistry(commands_server.ServerApi, commands_server.ServerAPICmd)
}
