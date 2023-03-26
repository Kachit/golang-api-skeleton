package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/lajosbencz/glo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MiddlewareHttpErrorHandlerMiddlewareTestSuite struct {
	suite.Suite
	logger   *infrastructure.LoggerMock
	gin      *gin.Context
	testable gin.HandlerFunc
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) SetupTest() {
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	suite.gin = c
	suite.logger = infrastructure.GetLoggerMock()
	suite.testable = HttpErrorHandlerMiddleware(suite.logger)
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) TestSuccess() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)

	ct := suite.gin.Writer.Header().Get("Content-Type")
	suite.testable(suite.gin)
	assert.Equal(suite.T(), http.StatusOK, suite.gin.Writer.Status())
	assert.NotEqual(suite.T(), ct, suite.gin.Writer.Header().Get("Content-Type"))
	assert.Equal(suite.T(), "application/json", suite.gin.Writer.Header().Get("Content-Type"))
	assert.Empty(suite.T(), suite.logger.Level)
	assert.Empty(suite.T(), suite.logger.Msg)
	assert.Empty(suite.T(), suite.logger.Params)
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) TestError() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)
	ct := suite.gin.Writer.Header().Get("Content-Type")
	suite.gin.AbortWithError(http.StatusForbidden, errors.New("foo bar"))
	suite.testable(suite.gin)
	assert.Equal(suite.T(), http.StatusForbidden, suite.gin.Writer.Status())
	assert.NotEqual(suite.T(), ct, suite.gin.Writer.Header().Get("Content-Type"))
	assert.Equal(suite.T(), "application/json", suite.gin.Writer.Header().Get("Content-Type"))
	assert.Equal(suite.T(), glo.Error, suite.logger.Level)
	assert.Equal(suite.T(), "foo bar", suite.logger.Msg)
	assert.NotEmpty(suite.T(), suite.logger.Params)
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) TestConvertErrorToHttpError() {
	err := errors.New("UsersApiResource.Create: UsersService.Create: user is not equal")
	result := convertErrorToHttpError(err)
	assert.Equal(suite.T(), "user is not equal", result.Error())
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) TestConvertErrorMessageWithColon() {
	msg := "UsersApiResource.Open: UsersService.Create: user is not equal"
	result := convertErrorMessage(msg)
	assert.Equal(suite.T(), "user is not equal", result)
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) TestConvertErrorMessageWithoutColon() {
	msg := "user is not equal"
	result := convertErrorMessage(msg)
	assert.Equal(suite.T(), "user is not equal", result)
}

func (suite *MiddlewareHttpErrorHandlerMiddlewareTestSuite) TestConvertErrorMessageSqlError() {
	msg := `UsersApiResource.Open: UsersService.Open: duplicate key value violates unique constraint "ux_user_email" (SQLSTATE 23505)`
	result := convertErrorMessage(msg)
	assert.Equal(suite.T(), "Database error", result)
}

func TestMiddlewareHttpErrorHandlerMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareHttpErrorHandlerMiddlewareTestSuite))
}
