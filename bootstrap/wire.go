//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/api"
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

func InitializePasswordGenerator(cfg *config.Config) (infrastructure.PasswordGenerator, error) {
	wire.Build(infrastructure.NewPasswordGenerator)
	return &infrastructure.PasswordGeneratorBCrypt{}, nil
}

func InitializeHashIds(cfg *config.Config) (*infrastructure.HashIds, error) {
	wire.Build(infrastructure.NewHashIds)
	return &infrastructure.HashIds{}, nil
}

func InitializeFractalManager() (*fractal.Manager, error) {
	wire.Build(infrastructure.NewFractalManager)
	return &fractal.Manager{}, nil
}

func InitializeRepositoriesFactory(db *gorm.DB) (*models.RepositoriesFactory, error) {
	wire.Build(models.NewRepositoriesFactory)
	return &models.RepositoriesFactory{}, nil
}

func InitializeMiddlewareFactory(container *infrastructure.Container) (*middleware.Factory, error) {
	wire.Build(middleware.NewMiddlewareFactory)
	return &middleware.Factory{}, nil
}

func InitializeErrorsResource(container *infrastructure.Container) (*api.ErrorsResource, error) {
	wire.Build(api.NewErrorsResource)
	return &api.ErrorsResource{}, nil
}

func InitializeDocumentationResource(container *infrastructure.Container) (*api.DocumentationResource, error) {
	return &api.DocumentationResource{}, nil
}

func InitializeUsersAPIResource(container *infrastructure.Container) (*api.UsersAPIResource, error) {
	wire.Build(api.NewUsersAPIResource)
	return &api.UsersAPIResource{}, nil
}

func InitializeContainer(configPath string) (*infrastructure.Container, error) {
	wire.Build(
		InitializeConfig,
		InitializeLogger,
		InitializePasswordGenerator,
		InitializeHashIds,
		InitializeFractalManager,
		InitializeDatabase,
		InitializeRepositoriesFactory,
		wire.Struct(new(infrastructure.Container),
			"Config",
			"Logger",
			"PG",
			"HashIds",
			"Fractal",
			"DB",
			"RF",
		),
	)
	return &infrastructure.Container{}, nil
}
