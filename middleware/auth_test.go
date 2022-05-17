package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Middleware_TokenAuthMiddlewareValidWithEnabledAuthAndV1UrlPath(t *testing.T) {
	token := "foo"
	header := "X-Auth-Token"
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/boxes", nil)
	c.Request.Header.Add(header, token)
	cfg := &config.AuthConfig{Enabled: true, Header: header, Token: token}
	mdl := TokenAuthMiddleware(cfg, nil)
	mdl(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func Test_Middleware_TokenAuthMiddlewareValidWithEnabledAuthAndNonV1UrlPathAndEmptyToken(t *testing.T) {
	token := "foo"
	header := "X-Auth-Token"
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/health-check", nil)
	cfg := &config.AuthConfig{Enabled: true, Header: header, Token: token}
	mdl := TokenAuthMiddleware(cfg, nil)
	mdl(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func Test_Middleware_TokenAuthMiddlewareValidWithDisabledAuthAndEmptyToken(t *testing.T) {
	token := "foo"
	header := "X-Auth-Token"
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/boxes", nil)
	cfg := &config.AuthConfig{Enabled: false, Header: header, Token: token}
	mdl := TokenAuthMiddleware(cfg, nil)
	mdl(c)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func Test_Middleware_TokenAuthMiddlewareInvalidWithEmptyToken(t *testing.T) {
	token := "foo"
	header := "X-Auth-Token"
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/boxes", nil)
	cfg := &config.AuthConfig{Enabled: true, Header: header, Token: token}
	mdl := TokenAuthMiddleware(cfg, nil)
	mdl(c)
	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
}

func Test_Middleware_TokenAuthMiddlewareInvalidWithWrongToken(t *testing.T) {
	token := "foo"
	header := "X-Auth-Token"
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/boxes", nil)
	c.Request.Header.Add(header, token+token)
	cfg := &config.AuthConfig{Enabled: true, Header: header, Token: token}
	mdl := TokenAuthMiddleware(cfg, nil)
	mdl(c)
	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
}

func Test_Middleware_TokenAuthMiddlewareInvalidWithEmptyTokenInConfig(t *testing.T) {
	token := "foo"
	header := "X-Auth-Token"
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "http://foo.bar/v1/boxes", nil)
	c.Request.Header.Add(header, token)
	cfg := &config.AuthConfig{Enabled: true}
	mdl := TokenAuthMiddleware(cfg, nil)
	mdl(c)
	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())
}
