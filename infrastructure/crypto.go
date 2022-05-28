package infrastructure

import (
	"github.com/kachit/golang-api-skeleton/config"
	"golang.org/x/crypto/bcrypt"
)

type PasswordGenerator interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) bool
}

func NewPasswordGenerator(config *config.Config) PasswordGenerator {
	return &PasswordGeneratorBCrypt{config.Crypto.Cost}
}

type PasswordGeneratorBCrypt struct {
	cost int
}

func (pg *PasswordGeneratorBCrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), pg.cost)
	return string(bytes), err
}

func (pg *PasswordGeneratorBCrypt) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
