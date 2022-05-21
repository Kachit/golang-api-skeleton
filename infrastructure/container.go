package infrastructure

import (
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/models"
	"gorm.io/gorm"
)

type Container struct {
	Config *config.Config
	Logger Logger
	DB     *gorm.DB
	PG     PasswordGenerator
	RF     *models.RepositoriesFactory
}
