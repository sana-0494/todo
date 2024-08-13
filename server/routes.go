package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/configs"

	"github.com/gin-gonic/gin"
)

func addTodoRoutes(group *gin.RouterGroup) {

	group.POST("/", func(ctx *gin.Context) {
	})
}

func Serve(ctx context.Context, cfg configs.Config) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Ok"})
	})
	todoGroup := router.Group("/todo")
	addTodoRoutes(todoGroup)

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutDownCtx, shutDownRelease := context.WithTimeout(ctx, 10*time.Second)
	defer shutDownRelease()

	if err := server.Shutdown(shutDownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Gracefull shutdown complete")
}
