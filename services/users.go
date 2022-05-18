package services

import (
	"encoding/json"
	"fmt"
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
)

func NewUsersService(container *infrastructure.Container) *UsersService {
	return &UsersService{rf: container.RF}
}

type UsersService struct {
	rf *models.RepositoriesFactory
}

func (us *UsersService) GetListByFilter() ([]*models.User, error) {
	rep := us.rf.GetUsersRepository()
	users, err := rep.GetListByFilter()
	if err != nil {
		return nil, fmt.Errorf("UsersService.GetListByFilter: %v", err)
	}
	return users, nil
}

func (us *UsersService) GetById(id uint64) (*models.User, error) {
	rep := us.rf.GetUsersRepository()
	user, err := rep.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("UsersService.GetById: %v", err)
	}
	return user, nil
}

func (us *UsersService) GetByEmail(email string) (*models.User, error) {
	rep := us.rf.GetUsersRepository()
	user, err := rep.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("UsersService.GetByEmail: %v", err)
	}
	return user, nil
}

func (us *UsersService) Create(userDto *dto.CreateUserDTO) (*models.User, error) {
	rep := us.rf.GetUsersRepository()
	err := us.checkIsUniqueEmail(userDto.Email, nil)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Create: %v", err)
	}
	user, err := us.buildUserFromCreateUserDTO(userDto)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Create: %v", err)
	}

	err = rep.Create(user)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Create: %v", err)
	}
	return user, nil
}

func (us *UsersService) Edit(id uint64, userDto *dto.EditUserDTO) (*models.User, error) {
	rep := us.rf.GetUsersRepository()
	user, err := rep.GetById(id)
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
	err = rep.Edit(user)
	if err != nil {
		return nil, fmt.Errorf("UsersService.Edit: %v", err)
	}
	return user, nil
}

func (us *UsersService) checkIsUniqueEmail(email string, user *models.User) error {
	rep := us.rf.GetUsersRepository()
	if user == nil || (email != "" && user.Email != email) {
		cnt, err := rep.CountByEmail(email)
		if err != nil {
			return err
		}
		if cnt > 0 {
			return fmt.Errorf("not unique user email")
		}
	}
	return nil
}

func (us *UsersService) buildUserFromCreateUserDTO(userDto *dto.CreateUserDTO) (*models.User, error) {
	data, err := json.Marshal(userDto)
	if err != nil {
		return nil, fmt.Errorf("UsersService.buildUserFromCreateUserDTO: %v", err)
	}
	var box models.User
	err = json.Unmarshal(data, &box)
	if err != nil {
		return nil, fmt.Errorf("UsersService.buildUserFromCreateUserDTO: %v", err)
	}
	return &box, nil
}

func (us *UsersService) fillUserFromEditUserDTO(user *models.User, userDto *dto.EditUserDTO) (*models.User, error) {
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
