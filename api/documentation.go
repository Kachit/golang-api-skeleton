package api

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//go:embed swagger.yml
var swaggerFile embed.FS

type DocumentationResource struct {
}

func NewDocumentationResource() *DocumentationResource {
	return &DocumentationResource{}
}

func (a *DocumentationResource) GetSwagger(c *gin.Context) {
	file, err := swaggerFile.ReadFile("swagger.yml")
	if err != nil {
		c.JSON(http.StatusBadRequest, NewResponseBodyError(err))
		return
	}
	content := strings.Replace(string(file), "%host-name%", c.Request.Host, 1)
	c.String(http.StatusOK, content)
}
