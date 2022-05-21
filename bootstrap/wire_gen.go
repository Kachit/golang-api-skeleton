// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/kachit/golang-api-skeleton/api"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/middleware"
	"github.com/kachit/golang-api-skeleton/models"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeConfig(configPath string) (*config.Config, error) {
	configConfig, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	return configConfig, nil
}

func InitializeLogger(cfg *config.Config) (infrastructure.Logger, error) {
	logger := infrastructure.NewLogger(cfg)
	return logger, nil
}

func InitializeDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := infrastructure.NewDatabase(cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitializePasswordGenerator(cfg *config.Config) (infrastructure.PasswordGenerator, error) {
	passwordGenerator := infrastructure.NewPasswordGenerator(cfg)
	return passwordGenerator, nil
}

func InitializeRepositoriesFactory(db *gorm.DB) (*models.RepositoriesFactory, error) {
	repositoriesFactory := models.NewRepositoriesFactory(db)
	return repositoriesFactory, nil
}

func InitializeMiddlewareFactory(container *infrastructure.Container) (*middleware.Factory, error) {
	factory := middleware.NewMiddlewareFactory(container)
	return factory, nil
}

func InitializeErrorsResource(container *infrastructure.Container) (*api.ErrorsResource, error) {
	errorsResource := api.NewErrorsResource(container)
	return errorsResource, nil
}

func InitializeUsersAPIResource(container *infrastructure.Container) (*api.UsersAPIResource, error) {
	usersAPIResource := api.NewUsersAPIResource(container)
	return usersAPIResource, nil
}

func InitializeContainer(configPath string) (*infrastructure.Container, error) {
	configConfig, err := InitializeConfig(configPath)
	if err != nil {
		return nil, err
	}
	logger, err := InitializeLogger(configConfig)
	if err != nil {
		return nil, err
	}
	passwordGenerator, err := InitializePasswordGenerator(configConfig)
	if err != nil {
		return nil, err
	}
	db, err := InitializeDatabase(configConfig)
	if err != nil {
		return nil, err
	}
	repositoriesFactory, err := InitializeRepositoriesFactory(db)
	if err != nil {
		return nil, err
	}
	container := &infrastructure.Container{
		Config: configConfig,
		Logger: logger,
		PG:     passwordGenerator,
		DB:     db,
		RF:     repositoriesFactory,
	}
	return container, nil
}

// wire.go:

func InitializeDocumentationResource(container *infrastructure.Container) (*api.DocumentationResource, error) {
	return &api.DocumentationResource{}, nil
}
