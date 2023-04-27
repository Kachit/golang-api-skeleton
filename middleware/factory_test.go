package middleware

import (
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MiddlewareFactorySuite struct {
	suite.Suite
	cfg      *infrastructure.Config
	testable *Factory
}

func (suite *MiddlewareFactorySuite) SetupTest() {
	suite.cfg = &infrastructure.Config{Auth: infrastructure.AuthConfig{Header: "X-Auth-Token"}}
	suite.testable = NewMiddlewareFactory(&infrastructure.Container{Config: suite.cfg})
}

func (suite *MiddlewareFactorySuite) TestBuildTokenAuthMiddleware() {
	result := suite.testable.BuildTokenAuthMiddleware()
	assert.NotEmpty(suite.T(), result)
}

func (suite *MiddlewareFactorySuite) TestBuildCorsMiddleware() {
	result := suite.testable.BuildCorsMiddleware()
	assert.NotEmpty(suite.T(), result)
}

func (suite *MiddlewareFactorySuite) TestBuildHttpErrorHandlerMiddleware() {
	result := suite.testable.BuildHttpErrorHandlerMiddleware()
	assert.NotEmpty(suite.T(), result)
}

func (suite *MiddlewareFactorySuite) TestBuildCorsMiddlewareConfig() {
	result := suite.testable.buildCorsMiddlewareConfig()
	assert.True(suite.T(), result.AllowCredentials)
	assert.True(suite.T(), result.AllowAllOrigins)
	assert.Equal(suite.T(), []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}, result.AllowMethods)
	assert.Equal(suite.T(), []string{"Origin", "Content-Length", "Content-Type", "Accept", "X-Requested-With", "X-Auth-Token"}, result.AllowHeaders)
}

func TestMiddlewareFactorySuite(t *testing.T) {
	suite.Run(t, new(MiddlewareFactorySuite))
}
