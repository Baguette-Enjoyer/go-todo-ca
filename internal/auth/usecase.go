package auth

import (
	"context"
	"baguette/go-todo-c/internal/models"
)

type UseCase interface {
	SignUp(ctx context.Context,email string,password string) (*models.User,error)
	SignIn(ctx context.Context,email string,password string) (string,error)
	GetUsers(ctx context.Context,page int,rpp int) ([]*models.User,error)
	GetUserById(ctx context.Context,user_id string)(*models.User,error)
}