package auth

import (
	"context"
	"baguette/go-todo-c/internal/models"
)

type Repository interface {
	FindByEmail(context context.Context,email string) (*models.User,error)
	SignUp(context context.Context,user *models.User) (error)
	GetUsers(ctx context.Context,page int,rpp int) ([]*models.User,error)
	GetUserById(ctx context.Context,user_id string)(*models.User,error)
}