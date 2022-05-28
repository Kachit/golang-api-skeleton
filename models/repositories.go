package models

import (
	"fmt"
	"gorm.io/gorm"
)

func NewRepositoriesFactory(database *gorm.DB) *RepositoriesFactory {
	return &RepositoriesFactory{db: database}
}

type RepositoriesFactory struct {
	db *gorm.DB
}

func (rf *RepositoriesFactory) GetUsersRepository() *UsersRepository {
	return &UsersRepository{db: rf.db}
}

type UsersRepository struct {
	db *gorm.DB
}

func (r *UsersRepository) Create(entity *User) error {
	err := r.db.Create(entity).Error
	if err != nil {
		return fmt.Errorf("UsersRepository.Create: %v", err)
	}
	return nil
}

func (r *UsersRepository) Edit(entity *User) error {
	err := r.db.Unscoped().Save(entity).Error
	if err != nil {
		return fmt.Errorf("UsersRepository.Edit: %v", err)
	}
	return nil
}

func (r *UsersRepository) GetListByFilter() ([]*User, error) {
	var records []*User
	tx := r.db.Unscoped()
	result := tx.Find(&records)
	if result.Error != nil {
		return nil, fmt.Errorf("UsersRepository.GetListByFilter: %v", result.Error)
	}
	return records, nil
}

func (r *UsersRepository) CountByFilter() (int64, error) {
	var count int64
	result := r.db.Model(&User{}).Unscoped().Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("UsersRepository.CountByFilter: %v", result.Error)
	}
	return count, nil
}

func (r *UsersRepository) GetById(id uint64) (*User, error) {
	var record User
	result := r.db.Unscoped().Find(&record, id)
	if result.Error != nil {
		return nil, fmt.Errorf("UsersRepository.GetById: %v", result.Error)
	}
	if record.Id == 0 {
		return nil, fmt.Errorf("UsersRepository.GetById: record not found")
	}
	return &record, nil
}

func (r *UsersRepository) GetByEmail(email string) (*User, error) {
	var record User
	result := r.db.Find(&record, "email = ?", email)
	if result.Error != nil {
		return nil, fmt.Errorf("UsersRepository.GetByEmail: %v", result.Error)
	}
	if record.Id == 0 {
		return nil, fmt.Errorf("UsersRepository.GetByEmail: record not found")
	}
	return &record, nil
}

func (r *UsersRepository) CountByEmail(email string) (int64, error) {
	var count int64
	result := r.db.Model(&User{}).Unscoped().Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("UsersRepository.CountByEmail: %v", result.Error)
	}
	return count, nil
}
