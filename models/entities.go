package models

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

type UsersCollection []*User

type User struct {
	Id          uint64      `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	Email       string      `json:"email" db:"email"`
	Password    string      `json:"password" db:"password"`
	Description null.String `json:"description" db:"description"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
}

func (p *User) BeforeInsert() {
	p.CreatedAt = time.Now()
}
