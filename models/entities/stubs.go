package entities

import (
	"github.com/jaswdr/faker"
	"time"
)

func NewUserEntityStub(fkr *faker.Faker) *User {
	if fkr == nil {
		f := faker.New()
		fkr = &f
	}

	entity := &User{
		Id:        uint64(fkr.RandomDigitNotNull()),
		Name:      fkr.Person().Name(),
		Email:     fkr.Internet().Email(),
		Password:  fkr.Hash().MD5(),
		CreatedAt: time.Now().UTC(),
	}
	return entity
}
