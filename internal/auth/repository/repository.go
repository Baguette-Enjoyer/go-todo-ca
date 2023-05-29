package repository

import (
	"context"
	"baguette/go-todo-c/internal/models"
	"baguette/go-todo-c/internal/auth"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository{
	return &authRepo{
		db:db,
	}
}

func (r *authRepo) FindByEmail(context context.Context,email string) (*models.User,error){
		u := models.User{}
		err := r.db.Table("users").Where("email = ?",email).First(&u).Error
		if err != nil {
			return nil,err
		}
		return &u,nil

}
func (r *authRepo) SignUp(context context.Context,user *models.User) (error){
	if err := r.db.Table("users").Create(&user).Error;err != nil {
		return err
	}
	return nil
}

func (r *authRepo) GetUsers(ctx context.Context,page int,rpp int) ([]*models.User,error){
	us := []*models.User{}
	startIndex := (page-1)*rpp
	err := r.db.Table("users").Find(&us).Limit(rpp).Offset(startIndex).Error
	if err != nil {
		return nil,err
	}
	return us,nil
}

func (r *authRepo) GetUserById(ctx context.Context,user_id string)(*models.User,error){
	user := &models.User{}
	err := r.db.Table("users").Where("user_id = ?",user_id).First(&user).Error
	if err != nil {
		return nil,err
	}
	return user,nil
}