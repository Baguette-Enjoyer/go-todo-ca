package repository

import (
	"context"
	"errors"

	"baguette/go-todo-c/internal/models"
	"baguette/go-todo-c/internal/todo"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.Repository{
	return &todoRepository{
		db:db,
	}
}

func (r *todoRepository) GetTodoById(context context.Context,todo_id string) (*models.Todo,error){
	todo := &models.Todo{}
	err := r.db.Table("todos").Where("todo_id = ?",todo_id).First(&todo).Error
	if err != nil {
		return nil,err
	}
	return todo,nil
}

func (r *todoRepository)AddTodo(context context.Context,todo *models.Todo)(error){
	err := r.db.Table("todos").Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *todoRepository)CheckTodo(context context.Context,todo_id string) (error) {
	todo,_ := r.GetTodoById(context,todo_id)
	if todo == nil {
		return errors.New("todo not found")
	}
	updateTodo := &models.Todo{
		TodoID: todo.TodoID,
		Done: !todo.Done,
	}
	err := r.db.Table("todos").Save(&updateTodo).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *todoRepository)GetTodos(context context.Context,user_id string) ([]*models.Todo,error) {
	todos := []*models.Todo{}
	err := r.db.Table("todos").Where("user_id = ?",user_id).Find(&todos).Error
	if err != nil {
		return nil,err
	}
	return todos,nil
}