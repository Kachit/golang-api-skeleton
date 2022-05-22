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
	us      *services.UsersService
	tf      *transformers.Factory
	log     infrastructure.Logger
	hashIds *infrastructure.HashIds
}

func NewUsersAPIResource(container *infrastructure.Container) *UsersAPIResource {
	return &UsersAPIResource{
		us:      services.NewUsersService(container),
		log:     container.Logger,
		tf:      transformers.NewTransformersFactory(container.Fractal, container.HashIds),
		hashIds: container.HashIds,
	}
}

func (a *UsersAPIResource) GetList(c *gin.Context) {
	collection, err := a.us.GetListByFilter()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.GetList: %v", err))
		return
	}
	data, err := a.tf.MapUsersResourceCollection(collection)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.GetList: %v", err))
		return
	}
	body := rest.NewResponseBodyWithPagination(data, 0, len(collection))
	c.JSON(http.StatusOK, body)
}

func (a *UsersAPIResource) GetById(c *gin.Context) {
	params := dto.HashIdUriParameterDTO{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.GetById: %v", err))
		return
	}
	id, err := a.hashIds.DecodeUint64(params.ID)
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.GetById: %v", err))
		return
	}
	user, err := a.us.GetById(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("UsersAPIResource.GetById: %v", err))
		return
	}

	data, err := a.tf.MapUsersResourceItem(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.GetById: %v", err))
		return
	}
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
	data, err := a.tf.MapUsersResourceItem(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Create: %v", err))
		return
	}
	body := rest.NewResponseBody(data)
	c.JSON(http.StatusCreated, body)
}

func (a *UsersAPIResource) Edit(c *gin.Context) {
	params := dto.HashIdUriParameterDTO{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	id, err := a.hashIds.DecodeUint64(params.ID)
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	userDTO, err := dto.BindEditUserDTO(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	user, err := a.us.Edit(id, userDTO)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	data, err := a.tf.MapUsersResourceItem(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("UsersAPIResource.Edit: %v", err))
		return
	}
	body := rest.NewResponseBody(data)
	c.JSON(http.StatusOK, body)
}
