package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/lajosbencz/glo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MiddlewareHttpErrorHandlerMiddlewareTestSuite struct {
	suite.Suite
	logger infrastructure.Logger
	gin    *gin.Context
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) SetupTest() {
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	suite.gin = c
	suite.logger = testable.GetLoggerMock()
}

func TestMiddlewareHttpErrorHandlerMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareHttpErrorHandlerMiddlewareTestSuite))
}

func Test_Middleware_HttpErrorHandlerMiddlewareSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)

	ct := c.Writer.Header().Get("Content-Type")
	logger := testable.GetLoggerMock()
	mdl := HttpErrorHandlerMiddleware(logger)
	mdl(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.NotEqual(t, ct, c.Writer.Header().Get("Content-Type"))
	assert.Equal(t, "application/json", c.Writer.Header().Get("Content-Type"))
	assert.Empty(t, logger.Level)
	assert.Empty(t, logger.Msg)
	assert.Empty(t, logger.Params)
}

func Test_Middleware_HttpErrorHandlerMiddlewareError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)
	ct := c.Writer.Header().Get("Content-Type")
	c.AbortWithError(http.StatusForbidden, errors.New("foo bar"))
	logger := testable.GetLoggerMock()
	mdl := HttpErrorHandlerMiddleware(logger)
	mdl(c)
	assert.Equal(t, http.StatusForbidden, c.Writer.Status())
	assert.NotEqual(t, ct, c.Writer.Header().Get("Content-Type"))
	assert.Equal(t, "application/json", c.Writer.Header().Get("Content-Type"))
	assert.Equal(t, glo.Error, logger.Level)
	assert.Equal(t, "foo bar", logger.Msg)
	assert.NotEmpty(t, logger.Params)
}

func Test_Middleware_ConvertErrorToHttpError(t *testing.T) {
	err := errors.New("UsersApiResource.Create: UsersService.Create: user is not equal")
	result := convertErrorToHttpError(err)
	assert.Equal(t, "user is not equal", result.Error())
}

func Test_Middleware_ConvertErrorMessageWithColon(t *testing.T) {
	msg := "UsersApiResource.Open: UsersService.Create: user is not equal"
	result := convertErrorMessage(msg)
	assert.Equal(t, "user is not equal", result)
}

func Test_Middleware_ConvertErrorMessageWithoutColon(t *testing.T) {
	msg := "user is not equal"
	result := convertErrorMessage(msg)
	assert.Equal(t, "user is not equal", result)
}

func Test_Middleware_ConvertErrorMessageSqlError(t *testing.T) {
	msg := `UsersApiResource.Open: UsersService.Open: duplicate key value violates unique constraint "ux_user_email" (SQLSTATE 23505)`
	result := convertErrorMessage(msg)
	assert.Equal(t, "Database error", result)
}
