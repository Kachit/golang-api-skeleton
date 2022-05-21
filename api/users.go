package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/dto"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/rest"
	"github.com/kachit/golang-api-skeleton/services"
	"github.com/kachit/golang-api-skeleton/transformers"
	"net/http"
)

type UsersAPIResource struct {
	us  *services.UsersService
	log infrastructure.Logger
}

func NewUsersAPIResource(container *infrastructure.Container) *UsersAPIResource {
	return &UsersAPIResource{us: services.NewUsersService(container), log: container.Logger}
}

func (a *UsersAPIResource) GetList(c *gin.Context) {
	collection, err := a.us.GetListByFilter()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.GetList: %v", err))
		return
	}
	data, err := transformers.MapUsersResourceCollection(collection)
	body := rest.NewResponseBodyWithPagination(data, 0, len(collection))
	c.JSON(http.StatusOK, body)
}

func (a *UsersAPIResource) GetById(c *gin.Context) {
	params := dto.IdUriParameterDTO{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.GetById: %v", err))
		return
	}
	user, err := a.us.GetById(params.ID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("UsersAPIResource.GetById: %v", err))
		return
	}

	data, _ := transformers.MapUsersResourceItem(user)
	c.JSON(http.StatusOK, rest.NewResponseBody(data))
}

func (a *UsersAPIResource) Create(c *gin.Context) {
	userDTO, err := dto.BindCreateUserDTO(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Create: %v", err))
		return
	}
	user, err := a.us.Create(userDTO)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Create: %v", err))
		return
	}
	body := rest.NewResponseBody(user)
	c.JSON(http.StatusCreated, body)
}

func (a *UsersAPIResource) Edit(c *gin.Context) {
	params := dto.IdUriParameterDTO{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	userDTO, err := dto.BindEditUserDTO(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	user, err := a.us.Edit(params.ID, userDTO)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	body := rest.NewResponseBody(user)
	c.JSON(http.StatusOK, body)
}
