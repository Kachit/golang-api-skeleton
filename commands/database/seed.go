package database

import (
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/kachit/golang-api-skeleton/database/seeders"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"github.com/urfave/cli/v2"
)

func NewDatabaseSeedCommand() *cli.Command {
	return &cli.Command{
		Name:  "database:seeders:seed",
		Usage: "Database seed dev data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yml",
				Usage: "Yml config file path",
			},
			&cli.UintFlag{
				Name:  "users",
				Usage: "Users rows count",
				Value: 10,
			},
		},
		Action: func(cCtx *cli.Context) error {
			configPath := cCtx.String("config")
			users := cCtx.Uint("users")
			container, err := bootstrap.InitializeContainer(configPath)
			if err != nil {
				return err
			}
			seedersStack := gorm_seeder.NewSeedersStack(container.DB)
			if users > 0 {
				cfg := gorm_seeder.SeederConfiguration{Rows: int(users)}
				usersSeeder := seeders.NewUsersSeeder(cfg, container.PG)
				seedersStack.AddSeeder(usersSeeder)
			}
			err = seedersStack.Seed()
			if err != nil {
				return err
			}
			return nil
		},
	}
}
