package dto

import (
	"github.com/jaswdr/faker"
)

func NewCreateUserDTOStub(fkr *faker.Faker) *CreateUserDTO {
	if fkr == nil {
		f := faker.New()
		fkr = &f
	}

	entity := &CreateUserDTO{
		Name:     fkr.Person().Name(),
		Email:    fkr.Internet().Email(),
		Password: fkr.Internet().Password(),
	}
	return entity
}

func NewEditUserDTOStub(fkr *faker.Faker) *EditUserDTO {
	if fkr == nil {
		f := faker.New()
		fkr = &f
	}

	entity := &EditUserDTO{
		Name:  fkr.Person().Name(),
		Email: fkr.Internet().Email(),
	}
	return entity
}
