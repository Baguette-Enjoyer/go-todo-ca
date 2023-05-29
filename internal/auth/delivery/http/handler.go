package http

import (
	"baguette/go-todo-c/internal/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	usecase auth.UseCase
}

func NewAuthHandler(usecase auth.UseCase) auth.Handler{
	return &authHandler{
		usecase: usecase,
	}
}

func (h *authHandler) SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {
		input := make(map[string]interface{})
		if err := c.Bind(&input);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err});
			return
		}
		// c.JSON(http.StatusOK,input)
		user,err := h.usecase.SignUp(c.Request.Context(),input["email"].(string),input["password"].(string))
		if err!= nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			return
		}
		c.JSON(http.StatusOK,user)
	}
}
func (h *authHandler) SignIn() gin.HandlerFunc{
	return func(c *gin.Context) {
		input := make(map[string]interface{})
		if err := c.Bind(&input);err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err});
			return
		}
		token,err := h.usecase.SignIn(c.Request.Context(),input["email"].(string),input["password"].(string))
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err});
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"token":token,
		})
	}
}

func (u *authHandler) GetUsers() gin.HandlerFunc{
	return func(c *gin.Context) {
		p := c.Query("page")
		page,_ :=  strconv.Atoi(p)
		if page == 0 {
			page = 1
		}
		rpp := 25
		rpp1 := c.Query("rpp")
		if rpp1 != ""{
			rpp,_ = strconv.Atoi(rpp1)
		}
		users,err := u.usecase.GetUsers(c,page,rpp)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err,
			})
			return
		}
		c.JSON(http.StatusOK,users)
	}
}