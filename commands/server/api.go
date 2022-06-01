package commands_server

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/mitchellh/cli"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func ServerAPICmd(ui cli.Ui) (cli.Command, error) {
	s := &ServerAPICommand{}
	s.init()
	return s, nil
}

type ServerAPICommand struct {
	flags      *flag.FlagSet
	configPath string
	helpText   string
}

func (s *ServerAPICommand) init() {
	s.flags = flag.NewFlagSet("", flag.ContinueOnError)
	s.flags.StringVar(&s.configPath, "config", "config.yml", "Yml config file path")
	//s.helpText = flags.Usage(s.help(), s.flags)
}

func (s *ServerAPICommand) help() string {
	return `
Usage: [options]

	API web server command
`
}

func (s *ServerAPICommand) Help() string {
	return s.helpText
}

func (s *ServerAPICommand) Synopsis() string {
	return ServerApi
}

func (s *ServerAPICommand) Run(args []string) int {
	if err := s.flags.Parse(args); err != nil {
		if !strings.Contains(err.Error(), "help requested") {
			log.Printf("error parsing flags: %v", err)
		}
		return 1
	}
	fmt.Println("Command launch " + s.Synopsis())
	err := s.RunServer()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}

func (s *ServerAPICommand) RunServer() error {
	container, err := bootstrap.InitializeContainer(s.configPath)
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
}
