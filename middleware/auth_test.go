package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MiddlewareTokenAuthMiddlewareTestSuite struct {
	suite.Suite
	cfg      *infrastructure.AuthConfig
	gin      *gin.Context
	testable gin.HandlerFunc
}

func (suite *MiddlewareTokenAuthMiddlewareTestSuite) SetupTest() {
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	suite.gin = c
	suite.cfg = &infrastructure.AuthConfig{Enabled: true, Header: "X-Auth-Token", Token: "foo"}
	suite.testable = TokenAuthMiddleware(suite.cfg, nil)
}

func (suite *MiddlewareTokenAuthMiddlewareTestSuite) TestValidWithEnabledAuthAndV1UrlPath() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)
	suite.gin.Request.Header.Add(suite.cfg.Header, suite.cfg.Token)
	suite.testable(suite.gin)
	assert.Equal(suite.T(), http.StatusOK, suite.gin.Writer.Status())
}

func (suite *MiddlewareTokenAuthMiddlewareTestSuite) TestInvalidWithEnabledAuthAndNonV1UrlPathAndEmptyToken() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/health-check", nil)
	suite.testable(suite.gin)
	assert.Equal(suite.T(), http.StatusUnauthorized, suite.gin.Writer.Status())
}

func (suite *MiddlewareTokenAuthMiddlewareTestSuite) TestValidWithDisabledAuthAndEmptyToken() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)
	suite.cfg.Enabled = false
	suite.testable(suite.gin)
	assert.Equal(suite.T(), http.StatusOK, suite.gin.Writer.Status())
}

func (suite *MiddlewareTokenAuthMiddlewareTestSuite) TestInvalidWithEmptyToken() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)
	suite.testable(suite.gin)
	assert.Equal(suite.T(), http.StatusUnauthorized, suite.gin.Writer.Status())
}

func (suite *MiddlewareTokenAuthMiddlewareTestSuite) TestInvalidWithWrongToken() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)
	suite.gin.Request.Header.Add(suite.cfg.Header, suite.cfg.Token+"qwerty")
	suite.testable(suite.gin)
	assert.Equal(suite.T(), http.StatusUnauthorized, suite.gin.Writer.Status())
}

func (suite *MiddlewareTokenAuthMiddlewareTestSuite) TestInvalidWithEmptyTokenInConfig() {
	suite.gin.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/users", nil)
	suite.gin.Request.Header.Add(suite.cfg.Header, suite.cfg.Token)
	suite.cfg.Token = ""
	mdl := TokenAuthMiddleware(suite.cfg, nil)
	mdl(suite.gin)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.gin.Writer.Status())
}

func TestMiddlewareTokenAuthMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTokenAuthMiddlewareTestSuite))
}
