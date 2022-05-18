//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/middleware"
	"github.com/kachit/golang-api-skeleton/models"
	"gorm.io/gorm"
)

func InitializeConfig(configPath string) (*config.Config, error) {
	wire.Build(config.NewConfig)
	return &config.Config{}, nil
}

func InitializeLogger(cfg *config.Config) (infrastructure.Logger, error) {
	wire.Build(infrastructure.NewLogger)
	return &infrastructure.LoggerAdapterGlo{}, nil
}

func InitializeDatabase(cfg *config.Config) (*gorm.DB, error) {
	wire.Build(infrastructure.NewDatabase)
	return &gorm.DB{}, nil
}

func InitializeRepositoriesFactory(db *gorm.DB) (*models.RepositoriesFactory, error) {
	wire.Build(models.NewRepositoriesFactory)
	return &models.RepositoriesFactory{}, nil
}

func InitializeMiddlewareFactory(container *infrastructure.Container) (*middleware.Factory, error) {
	wire.Build(middleware.NewMiddlewareFactory)
	return &middleware.Factory{}, nil
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
