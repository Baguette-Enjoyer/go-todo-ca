package todo

import "github.com/gin-gonic/gin"

type TodoHandler interface {
	GetTodoById() gin.HandlerFunc
	AddTodo() gin.HandlerFunc
	CheckTodo() gin.HandlerFunc
	GetTodos() gin.HandlerFunc
}
