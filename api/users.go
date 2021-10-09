package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/services"
	"net/http"
)

type UsersResource struct {
	UsersService *services.UsersService
}

func NewUsersResource(usersService *services.UsersService) *UsersResource {
	return &UsersResource{UsersService: usersService}
}

func (us *UsersResource) GetList(c *gin.Context) {
	var filter dto.FilterParameterQueryStringDTO
	if err := c.Bind(&filter); err != nil {
		c.JSON(http.StatusBadRequest, NewResponseBodyError(err))
		return
	}
	collection, err := us.UsersService.GetListByFilter(&filter)
	if err != nil {
		//fmt.Println(err)
		c.JSON(http.StatusBadRequest, NewResponseBodyError(err))
		return
	}
	body := NewResponseBodyCollection(collection, 1)
	c.JSON(http.StatusOK, body)
}

func (us *UsersResource) GetById(c *gin.Context) {
	userURI := dto.IdParameterPathDTO{}
	if err := c.ShouldBindUri(&userURI); err != nil {
		c.JSON(http.StatusBadRequest, NewResponseBodyError(err))
		return
	}
	user, err := us.UsersService.GetByID(userURI.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewResponseBodyError(err))
		return
	}
	body := NewResponseBody(user)
	c.JSON(http.StatusOK, body)
}
