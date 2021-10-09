package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, NewResponseBodyError(fmt.Errorf("Not found route")))
}

func NotAllowedMethodHandler(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, NewResponseBodyError(fmt.Errorf("Not allowed method")))
}
