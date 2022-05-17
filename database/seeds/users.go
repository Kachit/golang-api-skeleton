package seeds

import (
	"github.com/jaswdr/faker"
	"github.com/kachit/golang-api-skeleton/models"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type UsersSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewUsersSeeder(cfg gorm_seeder.SeederConfiguration) UsersSeeder {
	return UsersSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (s *UsersSeeder) Seed(db *gorm.DB) error {
	fkr := faker.New()
	var users []models.User
	for i := 0; i < s.Configuration.Rows; i++ {
		tag := models.User{
			Name:     fkr.Person().Name(),
			Email:    fkr.Internet().Email(),
			Password: fkr.Internet().Password(),
		}
		users = append(users, tag)
	}
	return db.CreateInBatches(users, s.Configuration.Rows).Error
}

func (s *UsersSeeder) Clear(db *gorm.DB) error {
	entity := models.User{}
	return s.SeederAbstract.Delete(db, entity.TableName())
}
