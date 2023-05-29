package server

import (
	authRepository "baguette/go-todo-c/internal/auth/repository"
	authUsecase "baguette/go-todo-c/internal/auth/usecase"
	authHttp "baguette/go-todo-c/internal/auth/delivery/http"
	
	todoRepository "baguette/go-todo-c/internal/todo/repository"
	todoUsecase "baguette/go-todo-c/internal/todo/usecase"
	todoHttp "baguette/go-todo-c/internal/todo/delivery/http"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandler(e *gin.Engine) error{
	//repo
	authRepo := authRepository.NewAuthRepository(s.db)
	todoRepo := todoRepository.NewTodoRepository(s.db)

	//usecase
	authUseCase := authUsecase.NewAuthUseCase(authRepo,[]byte(s.cfg.SigningKey))
	todoUsecase := todoUsecase.NewTodoUseCase(todoRepo,authRepo)
	//
	authHandler := authHttp.NewAuthHandler(authUseCase)
	todoHandler := todoHttp.NewTodoHandler(todoUsecase)
	//
	v1 := e.Group("/api/v1")
	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todos")

	//
	authHttp.MapAuthRoutes(authGroup,authHandler)
	todoHttp.MapTodosRoute(todoGroup,todoHandler)
	return nil
}