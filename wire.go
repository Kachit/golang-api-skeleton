//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/kachit/golang-api-skeleton/api"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
)

func InitializeConfig(configPath string) (*config.Config, error) {
	wire.Build(config.NewConfig)
	return &config.Config{}, nil
}

func InitializeDatabase(cfg *config.Config) (*sqlx.DB, error) {
	wire.Build(infrastructure.NewDatabase)
	return &sqlx.DB{}, nil
}

func InitializeLogger(cfg *config.Config) (infrastructure.Logger, error) {
	wire.Build(infrastructure.NewLogger)
	return &infrastructure.LoggerAdapterGlo{}, nil
}

func InitializeRepositoriesFactory(db *sqlx.DB) (*models.RepositoriesFactory, error) {
	wire.Build(models.NewRepositoriesFactory)
	return &models.RepositoriesFactory{}, nil
}

func InitializeUsersResource(container *infrastructure.Container) (*api.UsersResource, error) {
	wire.Build(api.NewUsersResource)
	return &api.UsersResource{}, nil
}

func InitializeErrorsResource(container *infrastructure.Container) (*api.ErrorsResource, error) {
	wire.Build(api.NewErrorsResource)
	return &api.ErrorsResource{}, nil
}

func InitializeContainer(configPath string) (*infrastructure.Container, error) {
	wire.Build(
		InitializeConfig,
		InitializeLogger,
		InitializeDatabase,
		InitializeRepositoriesFactory,
		wire.Struct(new(infrastructure.Container),
			"Config",
			"Logger",
			"DB",
			"RF",
		),
	)
	return &infrastructure.Container{}, nil
}
