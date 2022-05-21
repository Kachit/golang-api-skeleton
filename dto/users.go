package dto

import (
	"github.com/gin-gonic/gin"
)

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required" conform:"trim"`
	Email    string `json:"email" binding:"required,email" conform:"trim"`
	Password string `json:"password" binding:"required" conform:"trim"`
}

type EditUserDTO struct {
	Name     string `json:"name,omitempty" conform:"trim"`
	Email    string `json:"email,omitempty" binding:"email" conform:"trim"`
	Password string `json:"password,omitempty" conform:"trim"`
}

func BindCreateUserDTO(c *gin.Context) (*CreateUserDTO, error) {
	var userDTO CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		return nil, err
	}
	return &userDTO, nil
}

func BindEditUserDTO(c *gin.Context) (*EditUserDTO, error) {
	var userDTO EditUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		return nil, err
	}
	return &userDTO, nil
}
