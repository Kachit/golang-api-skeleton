package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	TableUsers string = "users"
)

// User
type User struct {
	Id         uint64         `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	CreatedAt  time.Time      `json:"created_at"`
	ModifiedAt *time.Time     `json:"modified_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now().UTC()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	t := time.Now().UTC()
	u.ModifiedAt = &t
	return
}

// TableName overrides
func (User) TableName() string {
	return TableUsers
}
