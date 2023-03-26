package seeders

import (
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models/entities"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type UsersSeeder struct {
	gorm_seeder.SeederAbstract
	pg infrastructure.PasswordGenerator
}

func NewUsersSeeder(cfg gorm_seeder.SeederConfiguration, pg infrastructure.PasswordGenerator) UsersSeeder {
	return UsersSeeder{SeederAbstract: gorm_seeder.NewSeederAbstract(cfg), pg: pg}
}

func (s *UsersSeeder) Seed(db *gorm.DB) error {
	fkr := faker.New()
	var users []entities.User
	for i := 0; i < s.Configuration.Rows; i++ {
		password, _ := s.pg.HashPassword(fmt.Sprintf("%d", 123456+i))
		user := entities.User{
			Name:     fkr.Person().Name(),
			Email:    fkr.Internet().Email(),
			Password: password,
		}
		users = append(users, user)
	}
	return db.CreateInBatches(users, s.Configuration.Rows).Error
}

func (s *UsersSeeder) Clear(db *gorm.DB) error {
	entity := entities.User{}
	return s.SeederAbstract.Delete(db, entity.TableName())
}
