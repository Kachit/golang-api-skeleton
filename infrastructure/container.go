package infrastructure

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/models"
	"gorm.io/gorm"
)

type Container struct {
	Config  *config.Config
	Logger  Logger
	HashIds *HashIds
	Fractal *fractal.Manager
	DB      *gorm.DB
	PG      PasswordGenerator
	RF      *models.RepositoriesFactory
}
