package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(config *config.AuthConfig, logger infrastructure.Logger) gin.HandlerFunc {
	requiredToken := config.Token
	tokenHeader := config.Header

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if config.Enabled && strings.Contains(path, "/v1/") {
			token := c.Request.Header.Get(tokenHeader)

			if requiredToken == "" || tokenHeader == "" {
				c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Auth error"))
				return
			}

			if token == "" {
				c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("API token required"))
				return
			}

			if token != requiredToken {
				c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Wrong API token"))
				return
			}
		}
		c.Next()
	}
}
