package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/models"
)

type UsersService struct {
	Database *sqlx.DB
}

func NewUsersService(database *sqlx.DB) *UsersService {
	return &UsersService{Database: database}
}

func (a *UsersService) GetListByFilter(filter *dto.FilterParameterQueryStringDTO) (*models.UsersCollection, error) {
	collection := models.UsersCollection{}
	return &collection, nil
}
