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
	c.AbortWithError(http.StatusNotFound, fmt.Errorf("Not found route"))
	return
}

func (er *ErrorsResource) NotAllowedMethodHandler(c *gin.Context) {
	c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("Not allowed method"))
	return
}
