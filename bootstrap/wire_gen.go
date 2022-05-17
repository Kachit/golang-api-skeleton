// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/infrastructure"
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

func InitializeRepositoriesFactory(db *gorm.DB) (*models.RepositoriesFactory, error) {
	repositoriesFactory := models.NewRepositoriesFactory(db)
	return repositoriesFactory, nil
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
		DB:     db,
		RF:     repositoriesFactory,
	}
	return container, nil
}
