package repositories

import (
	"fmt"
	"github.com/kachit/golang-api-skeleton/models/entities"
	"gorm.io/gorm"
)

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db}
}

type UsersRepository struct {
	db *gorm.DB
}

func (r *UsersRepository) Create(entity *entities.User) error {
	err := r.db.Create(entity).Error
	if err != nil {
		return fmt.Errorf("UsersRepository.Create: %v", err)
	}
	return nil
}

func (r *UsersRepository) Edit(entity *entities.User) error {
	err := r.db.Unscoped().Save(entity).Error
	if err != nil {
		return fmt.Errorf("UsersRepository.Edit: %v", err)
	}
	return nil
}

func (r *UsersRepository) GetListByFilter() ([]*entities.User, error) {
	var records []*entities.User
	tx := r.db.Unscoped()
	result := tx.Find(&records)
	if result.Error != nil {
		return nil, fmt.Errorf("UsersRepository.GetListByFilter: %v", result.Error)
	}
	return records, nil
}

func (r *UsersRepository) CountByFilter() (int64, error) {
	var count int64
	result := r.db.Model(&entities.User{}).Unscoped().Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("UsersRepository.CountByFilter: %v", result.Error)
	}
	return count, nil
}

func (r *UsersRepository) GetById(id uint64) (*entities.User, error) {
	var record entities.User
	result := r.db.Unscoped().Find(&record, id)
	if result.Error != nil {
		return nil, fmt.Errorf("UsersRepository.GetById: %v", result.Error)
	}
	if record.Id == 0 {
		return nil, fmt.Errorf("UsersRepository.GetById: record not found")
	}
	return &record, nil
}

func (r *UsersRepository) GetByEmail(email string) (*entities.User, error) {
	var record entities.User
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
	result := r.db.Model(&entities.User{}).Unscoped().Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("UsersRepository.CountByEmail: %v", result.Error)
	}
	return count, nil
}
