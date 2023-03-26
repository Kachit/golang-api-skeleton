// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/api"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/middleware"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeConfig(configPath string) (*infrastructure.Config, error) {
	config, err := infrastructure.NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func InitializeLogger(cfg *infrastructure.Config) (infrastructure.Logger, error) {
	logger := infrastructure.NewLogger(cfg)
	return logger, nil
}

func InitializeDatabase(cfg *infrastructure.Config) (*gorm.DB, error) {
	db, err := infrastructure.NewDatabase(cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitializePasswordGenerator(cfg *infrastructure.Config) (infrastructure.PasswordGenerator, error) {
	passwordGenerator := infrastructure.NewPasswordGenerator(cfg)
	return passwordGenerator, nil
}

func InitializeHashIds(cfg *infrastructure.Config) (*infrastructure.HashIds, error) {
	hashIds := infrastructure.NewHashIds(cfg)
	return hashIds, nil
}

func InitializeFractalManager() (*fractal.Manager, error) {
	manager := infrastructure.NewFractalManager()
	return manager, nil
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
	config, err := InitializeConfig(configPath)
	if err != nil {
		return nil, err
	}
	logger, err := InitializeLogger(config)
	if err != nil {
		return nil, err
	}
	passwordGenerator, err := InitializePasswordGenerator(config)
	if err != nil {
		return nil, err
	}
	hashIds, err := InitializeHashIds(config)
	if err != nil {
		return nil, err
	}
	manager, err := InitializeFractalManager()
	if err != nil {
		return nil, err
	}
	db, err := InitializeDatabase(config)
	if err != nil {
		return nil, err
	}
	container := &infrastructure.Container{
		Config:  config,
		Logger:  logger,
		PG:      passwordGenerator,
		HashIds: hashIds,
		Fractal: manager,
		DB:      db,
	}
	return container, nil
}

// wire.go:

func InitializeDocumentationResource(container *infrastructure.Container) (*api.DocumentationResource, error) {
	return &api.DocumentationResource{}, nil
}
