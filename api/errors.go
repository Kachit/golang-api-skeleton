package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"net/http"
)

func NewErrorsResource(container *infrastructure.Container) *ErrorsResource {
	return &ErrorsResource{
		logger: container.Logger,
	}
}

type ErrorsResource struct {
	logger infrastructure.Logger
}

func (er *ErrorsResource) NotFoundHandler(c *gin.Context) {
	err := "Not found route"
	er.logger.Warning(err)
	c.JSON(http.StatusNotFound, NewResponseBodyError(fmt.Errorf(err)))
}

func (er *ErrorsResource) NotAllowedMethodHandler(c *gin.Context) {
	err := "Not allowed method"
	er.logger.Warning(err)
	c.JSON(http.StatusMethodNotAllowed, NewResponseBodyError(fmt.Errorf(err)))
}
