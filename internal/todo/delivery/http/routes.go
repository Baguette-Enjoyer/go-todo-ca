package http

import (
	"baguette/go-todo-c/internal/todo"

	"github.com/gin-gonic/gin"
)

func MapTodosRoute(e *gin.RouterGroup,h todo.TodoHandler){
	e.GET("/:todo_id",h.GetTodoById())
	e.GET("/user/:user_id",h.GetTodos())
	e.POST("/",h.AddTodo())
	e.PUT("/:todo_id",h.CheckTodo())
}
