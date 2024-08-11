package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addTodoRoutes(group *gin.RouterGroup) {

	group.POST("/", func(ctx *gin.Context) {
		
	})
}


func Serve(ctx context.Context, logger log.Logger) {

	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Ok"})
	})

	todoGroup := router.Group("/todo")
	addTodoRoutes(todoGroup)

}
