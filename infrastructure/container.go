package infrastructure

import (
	"github.com/ibllex/go-fractal"
	"gorm.io/gorm"
)

type Container struct {
	Config  *Config
	Logger  Logger
	HashIds *HashIds
	Fractal *fractal.Manager
	DB      *gorm.DB
	PG      PasswordGenerator
}
