package http

import (
	"baguette/go-todo-c/internal/models"
	"baguette/go-todo-c/internal/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	usecase todo.TodoUseCase
}

func NewTodoHandler(usecase todo.TodoUseCase) todo.TodoHandler {
	return &todoHandler{
		usecase: usecase,
	}
}

func (h *todoHandler) GetTodoById() gin.HandlerFunc {
	return func(c *gin.Context) {
		todo_id := c.Param("todo_id")
		todo, err := h.usecase.GetTodoById(c, todo_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}
func (h *todoHandler) AddTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		todo := models.Todo{}
		if err := c.Bind(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		err := h.usecase.AddTodo(c, &todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})

	}
}
func (h *todoHandler) CheckTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		todo_id := c.Param("todo_id")
		err := h.usecase.CheckTodo(c, todo_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
func (h *todoHandler) GetTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		todos,err := h.usecase.GetTodos(c,user_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err,
			})
			return
		}
		c.JSON(http.StatusOK,todos)
	}
}
