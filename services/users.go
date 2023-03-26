package services

import (
	"encoding/json"
	"fmt"
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models/entities"
	"github.com/kachit/golang-api-skeleton/models/repositories"
)

func NewUsersService(container *infrastructure.Container) *UsersService {
	return &UsersService{
		ur: repositories.NewUsersRepository(container.DB),
		pg: container.PG,
	}
}

type UsersService struct {
	ur *repositories.UsersRepository
	pg infrastructure.PasswordGenerator
}

func (us *UsersService) GetListByFilter() ([]*entities.User, error) {
	users, err := us.ur.GetListByFilter()
	if err != nil {
		return nil, fmt.Errorf("UsersService.GetListByFilter: %v", err)
	}
	return users, nil
}

func (us *UsersService) CountByFilter() (int64, error) {
	cnt, err := us.ur.CountByFilter()
	if err != nil {
		return 0, fmt.Errorf("UsersService.CountByFilter: %v", err)
	}
	return cnt, nil
}

func (us *UsersService) GetById(id uint64) (*entities.User, error) {
	user, err := us.ur.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("UsersService.GetById: %v", err)
	}
	return user, nil
}

func (us *UsersService) Create(userDto *dto.CreateUserDTO) (*entities.User, error) {
	err := us.checkIsUniqueEmail(userDto.Email, nil)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Create: %v", err)
	}
	user, err := us.buildUserFromCreateUserDTO(userDto)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Create: %v", err)
	}
	user.Password, err = us.pg.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Create: %v", err)
	}
	err = us.ur.Create(user)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Create: %v", err)
	}
	return user, nil
}

func (us *UsersService) Edit(id uint64, userDto *dto.EditUserDTO) (*entities.User, error) {
	user, err := us.ur.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Edit: %v", err)
	}
	err = us.checkIsUniqueEmail(userDto.Email, user)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Edit: %v", err)
	}
	user, err = us.fillUserFromEditUserDTO(user, userDto)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Edit: %v", err)
	}
	err = us.ur.Edit(user)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Edit: %v", err)
	}
	return user, nil
}

func (us *UsersService) checkIsUniqueEmail(email string, user *entities.User) error {
	if user == nil || (email != "" && user.Email != email) {
		cnt, err := us.ur.CountByEmail(email)
		if err != nil {
			return err
		}
		if cnt > 0 {
			return fmt.Errorf("not unique user email")
		}
	}
	return nil
}

func (us *UsersService) buildUserFromCreateUserDTO(userDto *dto.CreateUserDTO) (*entities.User, error) {
	data, err := json.Marshal(userDto)
	if err != nil {
		return nil, fmt.Errorf("UsersService.buildUserFromCreateUserDTO: %v", err)
	}
	var user entities.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("UsersService.buildUserFromCreateUserDTO: %v", err)
	}
	return &user, nil
}

func (us *UsersService) fillUserFromEditUserDTO(user *entities.User, userDto *dto.EditUserDTO) (*entities.User, error) {
	data, err := json.Marshal(userDto)
	if err != nil {
		return nil, fmt.Errorf("UsersService.fillUserFromEditUserDTO: %v", err)
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("UsersService.fillUserFromEditUserDTO: %v", err)
	}
	return user, nil
}
