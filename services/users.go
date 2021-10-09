package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/kachit/golang-api-skeleton/dto"
)

type UsersService struct {
	Database *sqlx.DB
}

func NewUsersService(database *sqlx.DB) *UsersService {
	return &UsersService{Database: database}
}

func (a *UsersService) List() ([]dto.UserDTO, error) {
	var items = make([]dto.UserDTO, 0)

	return items, nil
}
