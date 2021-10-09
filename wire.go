//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/kachit/golang-api-skeleton/api"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/services"
)

func InitializeConfig(configPath string) (*config.Config, error) {
	wire.Build(config.NewConfig)
	return &config.Config{}, nil
}

func InitializeDatabase(cfg *config.Config) (*sqlx.DB, error) {
	wire.Build(infrastructure.NewDatabase)
	return &sqlx.DB{}, nil
}

func InitializeUsersResource(db *sqlx.DB) (*api.UsersResource, error) {
	wire.Build(services.NewUsersService, api.NewUsersResource)
	return &api.UsersResource{}, nil
}
