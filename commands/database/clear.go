package database

import (
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/kachit/golang-api-skeleton/database/seeders"
	gormseeder "github.com/kachit/gorm-seeder"
	"github.com/urfave/cli/v2"
)

func NewDatabaseClearCommand() *cli.Command {
	return &cli.Command{
		Name:  "database:seeders:clear",
		Usage: "Database clear dev data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yml",
				Usage: "Yml config file path",
			},
			&cli.BoolFlag{
				Name:  "confirm",
				Usage: "Confirmation flag",
				Value: false,
			},
		},
		Action: func(cCtx *cli.Context) error {
			configPath := cCtx.String("config")
			confirm := cCtx.Bool("confirm")
			container, err := bootstrap.InitializeContainer(configPath)
			if err != nil {
				return err
			}
			seedersStack := gormseeder.NewSeedersStack(container.DB)
			if confirm {
				cfg := gormseeder.SeederConfiguration{}
				usersSeeder := seeders.NewUsersSeeder(cfg, nil)
				seedersStack.AddSeeder(&usersSeeder)
			}
			err = seedersStack.Clear()
			if err != nil {
				return err
			}
			return nil
		},
	}
}
