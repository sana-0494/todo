package api

import (
	"context"
	"net/http"
	"time"
	td "todo/core"

	"github.com/gin-gonic/gin"
)

type TodoService interface {
	Create(ctx context.Context, td td.Todo) string
	List(ctx context.Context) td.Todo
	Update(ctx context.Context, id string, td td.Todo) td.Todo
	Delete(ctx context.Context, id string) bool
	Restore(ctx context.Context, id string) (td.Todo, error)
}

type TodoHandler struct {
	todoService TodoService
}

func NewHandler(t TodoService) TodoHandler {
	return TodoHandler{
		todoService: t,
	}
}

func (h TodoHandler) Create(ctx *gin.Context) {

	t := todoCreateRequest{}
	if err := ctx.ShouldBindBodyWithJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	id := h.todoService.Create(ctx, td.Todo{
		CreatedAt: time.Now(),
		Title:     t.Title,
		Status:    "completed",
	})
	//add err handling
	ctx.JSON(http.StatusOK, id)
}

func (h TodoHandler) List(ctx *gin.Context) {
	listTodo := h.todoService.List(ctx)
	ctx.JSON(http.StatusOK, listTodo)
}

func (h TodoHandler) Update(ctx *gin.Context) {
	todoId := ctx.Param("id")
	t := todoCreateRequest{}
	if err := ctx.ShouldBindBodyWithJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	updatedTodo := h.todoService.Update(ctx, todoId, td.Todo{
		CreatedAt: time.Now(),
		Title:     t.Title,
		Status:    "completed",
	})
	ctx.JSON(http.StatusOK, updatedTodo)
}

func (h TodoHandler) Delete(ctx *gin.Context) {
	todoId := ctx.Param("id")
	deletedTodo := h.todoService.Delete(ctx, todoId)
	if deletedTodo {
		ctx.JSON(http.StatusOK, todoDeleteResponse{
			Status:  "deleted",
			Message: "Your todo is temprorarily delete. It can be restored within 30 days of deletion",
		})
	}
}

func (h TodoHandler) Restore(ctx *gin.Context) {
	todoId := ctx.Param("id")
	RestoredTodo, err := h.todoService.Restore(ctx, todoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, RestoredTodo)
}
