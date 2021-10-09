package services

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/models"
)

type UsersService struct {
	repositories *models.RepositoriesFactory
}

func NewUsersService(database *sqlx.DB) *UsersService {
	return &UsersService{repositories: models.NewRepositoriesFactory(database)}
}

func (us *UsersService) GetListByFilter(filter *dto.FilterParameterQueryStringDTO) (*models.UsersCollection, error) {
	var condition interface{}
	collection, err := us.repositories.GetUsersRepository().GetList(condition, filter.Limit, filter.Offset, filter.Sort)
	if err != nil {
		return nil, fmt.Errorf("UsersService.GetListByFilter: %v", err)
	}
	return collection, nil
}

func (us *UsersService) GetByID(id uint64) (*models.User, error) {
	user, err := us.repositories.GetUsersRepository().GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("UsersService.GetByID: %v", err)
	}
	return user, nil
}
