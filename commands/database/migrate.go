package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/kachit/golang-api-skeleton/database/migrations"
	"github.com/urfave/cli/v2"
)

func NewDatabaseMigrateCommand() *cli.Command {
	return &cli.Command{
		Name:  "database:migrations:migrate",
		Usage: "Apply database migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yml",
				Usage: "Yml config file path",
			},
		},
		Action: func(cCtx *cli.Context) error {
			configPath := cCtx.String("config")
			container, err := bootstrap.InitializeContainer(configPath)
			if err != nil {
				return err
			}
			m := gormigrate.New(
				container.DB,
				gormigrate.DefaultOptions,
				migrations.Migrations,
			)
			err = m.Migrate()
			if err != nil {
				return err
			}
			return nil
		},
	}
}
