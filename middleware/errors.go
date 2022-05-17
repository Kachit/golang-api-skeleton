package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/rest"
	"github.com/kachit/golang-api-skeleton/utils"
	"strings"
)

func HttpErrorHandlerMiddleware(logger infrastructure.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		dump := utils.DumpHttpRequest(c.Request)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
		var err error
		if len(c.Errors) > 0 {
			err = c.Errors[0]
			logger.Error(err.Error(), dump)
			c.JSON(-1, rest.NewResponseBodyError(convertErrorToHttpError(err)))
		}
		return
	}
}

func convertErrorToHttpError(err error) error {
	return errors.New(convertErrorMessage(err.Error()))
}

func convertErrorMessage(msg string) string {
	var result string
	msgArr := strings.Split(msg, ":")
	result = strings.Trim(msgArr[len(msgArr)-1], " ")
	if strings.Contains(result, "SQLSTATE") {
		result = "Database error"
	}
	return result
}
