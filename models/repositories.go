package models

import "gorm.io/gorm"

func NewRepositoriesFactory(database *gorm.DB) *RepositoriesFactory {
	return &RepositoriesFactory{db: database}
}

type RepositoriesFactory struct {
	db *gorm.DB
}
