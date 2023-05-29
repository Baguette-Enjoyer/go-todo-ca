package usecase

import (
	"baguette/go-todo-c/internal/auth"
	"baguette/go-todo-c/internal/models"
	"baguette/go-todo-c/internal/todo"
	"context"
	// "errors"

	"github.com/google/uuid"
)

type todoUC struct {
	todoRepo todo.Repository
	userRepo auth.Repository
}

func NewTodoUseCase(todoRepo todo.Repository,userRepo auth.Repository) todo.TodoUseCase{
	return &todoUC{
		todoRepo: todoRepo,
		userRepo: userRepo,
	}
}

func (u *todoUC) GetTodoById(context context.Context,todo_id string) (*models.Todo,error){
	todo,err := u.todoRepo.GetTodoById(context,todo_id)
	if err != nil {
		return nil,err
	}
	return todo,nil
}

func (u *todoUC) AddTodo(context context.Context,todo *models.Todo) (error){
	insertTodo := &models.Todo{
		TodoID: uuid.New().String(),
		Content: todo.Content,
		UserID: todo.UserID,
	}
	err := u.todoRepo.AddTodo(context,insertTodo)
	if err != nil {
		return err
	}
	return nil
}

func (u *todoUC) CheckTodo(context context.Context,todo_id string) (error){
	err := u.todoRepo.CheckTodo(context,todo_id)
	if err != nil {
		return err
	}
	return nil
}

func (u *todoUC) GetTodos(context context.Context,user_id string) ([]*models.Todo,error){
	todos,err := u.todoRepo.GetTodos(context,user_id)
	if err != nil {
		return nil,err
	}
	return todos,nil
}