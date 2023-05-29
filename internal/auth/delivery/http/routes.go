package http

import (
	"baguette/go-todo-c/internal/auth"

	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroup *gin.RouterGroup,h auth.Handler){
	authGroup.GET("/users",h.GetUsers())
	authGroup.POST("/register",h.SignUp())
	authGroup.POST("/signin",h.SignIn())
}