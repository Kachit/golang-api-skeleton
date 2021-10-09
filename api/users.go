package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/services"
	"net/http"
)

type UsersResource struct {
	UsersService *services.UsersService
}

func NewUsersResource(usersService *services.UsersService) *UsersResource {
	return &UsersResource{UsersService: usersService}
}

func (a *UsersResource) GetList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
