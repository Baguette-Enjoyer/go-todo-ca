package todo

import (
	"baguette/go-todo-c/internal/models"
	"context"
)

type TodoUseCase interface {
	GetTodoById(context context.Context,todo_id string) (*models.Todo,error)
	AddTodo(context context.Context,todo *models.Todo) (error)
	CheckTodo(context context.Context,todo_id string) (error)
	GetTodos(context context.Context,user_id string) ([]*models.Todo,error)
	
}