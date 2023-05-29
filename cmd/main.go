package main

import (
	"log"
	"baguette/go-todo-c/config"
	"baguette/go-todo-c/db"
	"baguette/go-todo-c/internal/server"
	"github.com/sirupsen/logrus"
)

func main(){
	log.Println("Start server")
	cfg := config.NewConfig("../.env")
	db := db.GetPostgresInstance(cfg,true)
	s := server.NewServer(cfg,db,logrus.New())
	
	s.Run()
}