package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/kachit/golang-api-skeleton/database/migrations"
	"github.com/urfave/cli/v2"
)

func NewDatabaseRollbackCommand() *cli.Command {
	return &cli.Command{
		Name:  "database:migrations:rollback",
		Usage: "Rollback database migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yml",
				Usage: "Yml config file path",
			},
			&cli.StringFlag{
				Name:  "migration",
				Usage: "Migration id",
			},
		},
		Action: func(cCtx *cli.Context) error {
			configPath := cCtx.String("config")
			migration := cCtx.String("migration")

			container, err := bootstrap.InitializeContainer(configPath)
			if err != nil {
				return err
			}
			m := gormigrate.New(
				container.DB,
				gormigrate.DefaultOptions,
				migrations.Migrations,
			)
			if migration != "" {
				err = m.RollbackTo(migration)
			} else {
				err = m.RollbackLast()
			}
			if err != nil {
				return err
			}
			return nil
		},
	}
}
