package application

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewApplicationStartCommand() *cli.Command {
	return &cli.Command{
		Name:  "app:start",
		Usage: "Start web server command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yml",
				Usage: "Yml config file path",
			},
		},
		Action: func(cCtx *cli.Context) error {
			configPath := cCtx.String("config")
			container, err := bootstrap.InitializeContainer(configPath)
			if err != nil {
				return err
			}
			middlewareFactory, err := bootstrap.InitializeMiddlewareFactory(container)
			if err != nil {
				return err
			}

			errorsResource, err := bootstrap.InitializeErrorsResource(container)
			if err != nil {
				return err
			}

			docsApi, err := bootstrap.InitializeDocumentationResource(container)
			if err != nil {
				return err
			}

			usersApi, err := bootstrap.InitializeUsersAPIResource(container)
			if err != nil {
				return err
			}

			router := gin.Default()
			router.Use(middlewareFactory.BuildCorsMiddleware())
			router.Use(middlewareFactory.BuildHttpErrorHandlerMiddleware())
			router.NoRoute(errorsResource.NotFoundHandler)
			router.NoMethod(errorsResource.NotAllowedMethodHandler)
			//
			cfg := container.Config
			if !cfg.App.Debug {
				gin.SetMode(gin.ReleaseMode)
			}

			apiRoutes := router.Group("/")
			{
				shared := apiRoutes.Group("/shared")
				{
					shared.GET("/swagger", docsApi.GetSwagger)
				}
			}

			apiRoutesProtected := router.Group("/v1", middlewareFactory.BuildTokenAuthMiddleware())
			{
				users := apiRoutesProtected.Group("/users")
				{
					users.GET("", usersApi.GetList)
					users.GET("/:id", usersApi.GetById)
					users.POST("", usersApi.Create)
					users.PUT("/:id", usersApi.Edit)
				}
			}

			srv := &http.Server{
				Addr:    cfg.GetAppPort(),
				Handler: router,
			}
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			quit := make(chan os.Signal)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutting down server...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown:", err)
			}

			log.Println("Server exiting")
			return nil
		},
	}
}
