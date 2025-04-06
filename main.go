package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gin-boilerplate/config"
	"gin-boilerplate/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	router := gin.Default()
    api := router.Group("/api/v1")
	routes.RegisterUserRoutes(api)
	routes.AuthUserRoutes(api)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v\n", err)
		}
	}()
	log.Println("Server started on :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
	log.Println("Server exited properly")
}
