package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/models"
	"net/http"
)

type Container struct {
	Config *config.Config
	//Logger     Logger
	HttpClient *http.Client

	DB *sqlx.DB
	RF *models.RepositoriesFactory
}
