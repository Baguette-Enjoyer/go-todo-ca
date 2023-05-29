package usecase

import (
	"context"
	"errors"
	"time"

	"baguette/go-todo-c/internal/auth"
	"baguette/go-todo-c/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type userClaim struct {
	jwt.StandardClaims
	UserID	string
	Email 	string
}

type authUC struct {
	userRepo auth.Repository
	signingKey     []byte
}

func NewAuthUseCase(userRepo auth.Repository,signingKey []byte) auth.UseCase {
	return &authUC{
		userRepo: userRepo,
		signingKey: signingKey,
	}
}

func (u *authUC) SignUp(ctx context.Context, email string, password string) (*models.User, error) {
	existUser, _ := u.userRepo.FindByEmail(ctx, email)
	if existUser != nil {
		return nil, errors.New("user existed")
	}

	user := &models.User{
		UserID: uuid.New().String(),
		Email: email,
		Password: password,
	}
	user.HashPassword()
	err := u.userRepo.SignUp(ctx,user)
	if err != nil {
		return nil, err
	}
	return user,nil

}

func (u *authUC) SignIn(ctx context.Context, email string, password string) (string, error) {
	user,_ := u.userRepo.FindByEmail(ctx,email)
	if user == nil {
		return "",errors.New("user not exist")
	}
	if !user.ComparePassword(password) {
		return "",errors.New("wrong password")
	}
	claims := userClaim{
		UserID: user.UserID,
		Email:user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10*time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(u.signingKey)

}

func (u *authUC) GetUsers(ctx context.Context,page int,rpp int) ([]*models.User,error){
	users,err := u.userRepo.GetUsers(ctx,page,rpp)
	if err != nil {
		return []*models.User{},errors.New("error getting users list")
	}
	return users,nil
}

func (u *authUC) GetUserById(ctx context.Context,user_id string)(*models.User,error){
	user,err := u.userRepo.GetUserById(ctx,user_id)
	if err != nil {
		return nil,err
	}
	return user,nil
}
