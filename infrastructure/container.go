package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/models"
)

type Container struct {
	Config *config.Config
	Logger Logger

	DB *sqlx.DB
	RF *models.RepositoriesFactory
}
