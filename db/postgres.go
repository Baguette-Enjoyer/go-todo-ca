package db

import (
	"baguette/go-todo-c/config"
	"baguette/go-todo-c/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(cfg *config.Configuration,migrate bool) *gorm.DB {
	dsn := cfg.DBConnUrl
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if migrate {
		db.AutoMigrate(&models.Todo{},&models.User{})
	}
	return db
}