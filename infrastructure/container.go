package infrastructure

import (
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/models"
	"gorm.io/gorm"
	"net/http"
)

type Container struct {
	Config      *config.Config
	Logger      Logger
	HttpRequest http.Request
	DB          *gorm.DB
	RF          *models.RepositoriesFactory
}
