package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/infrastructure"
)

func NewMiddlewareFactory(container *infrastructure.Container) *Factory {
	return &Factory{container}
}

type Factory struct {
	cnt *infrastructure.Container
}

func (f *Factory) BuildTokenAuthMiddleware() gin.HandlerFunc {
	return TokenAuthMiddleware(&f.cnt.Config.Auth, f.cnt.Logger)
}

func (f *Factory) BuildHttpErrorHandlerMiddleware() gin.HandlerFunc {
	return HttpErrorHandlerMiddleware(f.cnt.Logger)
}

func (f *Factory) BuildCorsMiddleware() gin.HandlerFunc {
	config := f.buildCorsMiddlewareConfig()
	return cors.New(config)
}

func (f *Factory) buildCorsMiddlewareConfig() cors.Config {
	config := cors.DefaultConfig()
	headers := []string{"Accept", "X-Requested-With", f.cnt.Config.Auth.Header}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, headers...)
	return config
}
