package middleware

import (
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Middleware_BuildTokenAuthMiddleware(t *testing.T) {
	cfg := &config.Config{Auth: config.AuthConfig{}}
	cnt := &infrastructure.Container{Config: cfg}
	f := NewMiddlewareFactory(cnt)
	result := f.BuildTokenAuthMiddleware()
	assert.NotEmpty(t, result)
}

func Test_Middleware_BuildCorsMiddleware(t *testing.T) {
	cfg := &config.Config{Auth: config.AuthConfig{}}
	cnt := &infrastructure.Container{Config: cfg}
	f := NewMiddlewareFactory(cnt)
	result := f.BuildCorsMiddleware()
	assert.NotEmpty(t, result)
}

func Test_Middleware_BuildHttpErrorHandlerMiddleware(t *testing.T) {
	cfg := &config.Config{}
	cnt := &infrastructure.Container{Config: cfg}
	f := NewMiddlewareFactory(cnt)
	result := f.BuildHttpErrorHandlerMiddleware()
	assert.NotEmpty(t, result)
}

func Test_Middleware_BuildCorsMiddlewareConfig(t *testing.T) {
	cfg := &config.Config{Auth: config.AuthConfig{Header: "X-Auth-Token"}}
	cnt := &infrastructure.Container{Config: cfg}
	f := NewMiddlewareFactory(cnt)
	result := f.buildCorsMiddlewareConfig()
	assert.True(t, result.AllowCredentials)
	assert.True(t, result.AllowAllOrigins)
	assert.Equal(t, []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}, result.AllowMethods)
	assert.Equal(t, []string{"Origin", "Content-Length", "Content-Type", "Accept", "X-Requested-With", "X-Auth-Token"}, result.AllowHeaders)
}
