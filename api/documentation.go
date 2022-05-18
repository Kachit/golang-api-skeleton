package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

////go:embed swagger.yml
//var swaggerFile embed.FS

type DocumentationResource struct {
}

func NewDocumentationResource() *DocumentationResource {
	return &DocumentationResource{}
}

func (a *DocumentationResource) GetSwagger(c *gin.Context) {
	//file, err := swaggerFile.ReadFile("swagger.yml")
	//if err != nil {
	//	c.AbortWithError(http.StatusBadRequest, fmt.Errorf("DocumentationResource.GetSwagger: %v", err))
	//	return
	//}
	//content := strings.Replace(string(file), "%host-name%", c.Request.Host, 1)
	c.String(http.StatusOK, "")
}
