package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//configPathPtr := flag.String("config", "config.yml", "Path to configuration file")
	//flag.Parse()

	router := gin.Default()
	router.Use(cors.Default())
	//router.NoRoute(errorsResource.NotFoundHandler)
	//router.NoMethod(errorsResource.NotAllowedMethodHandler)

	//cfg := container.Config
	//if !cfg.App.Debug {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	//apiRoutes := router.Group("/v1")
	//{
	//	users := apiRoutes.Group("/users")
	//	{
	//		users.GET("", usersResource.GetList)
	//		users.GET("/:id", usersResource.GetById)
	//	}
	//}

	srv := &http.Server{
		//Addr:    cfg.GetAppPort(),
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
