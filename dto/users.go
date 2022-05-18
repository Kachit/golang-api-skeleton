package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/leebenson/conform"
)

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required" conform:"trim"`
	Email    string `json:"email" binding:"required" conform:"trim"`
	Password string `json:"password" binding:"required" conform:"trim"`
}

type EditUserDTO struct {
	Name     string `json:"name,omitempty" conform:"trim"`
	Email    string `json:"slug,omitempty" conform:"trim"`
	Password string `json:"image,omitempty" conform:"trim"`
}

func BindCreateUserDTO(c *gin.Context) (*CreateUserDTO, error) {
	var userDTO CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		return nil, err
	}
	if err := conform.Strings(&userDTO); err != nil {
		return nil, err
	}
	return &userDTO, nil
}

func BindEditUserDTO(c *gin.Context) (*EditUserDTO, error) {
	var userDTO EditUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		return nil, err
	}
	if err := conform.Strings(&userDTO); err != nil {
		return nil, err
	}
	return &userDTO, nil
}
