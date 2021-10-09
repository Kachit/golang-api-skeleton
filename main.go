package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kachit/golang-api-skeleton/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configPathPtr := flag.String("config", "config.yml", "Path to configuration file")
	flag.Parse()

	cfg, err := InitializeConfig(*configPathPtr)
	if err != nil {
		fmt.Println(err)
	}

	db, err := InitializeDatabase(cfg)
	if err != nil {
		fmt.Println(err)
	}

	usersApi, err := InitializeUsersResource(db)
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.NoRoute(api.NotFoundHandler)
	router.NoMethod(api.NotAllowedMethodHandler)

	apiRoutes := router.Group("/v1")
	{
		users := apiRoutes.Group("/users")
		{
			users.GET("", usersApi.GetList)
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
}
