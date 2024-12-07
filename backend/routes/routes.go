package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func signupRouter(c *gin.Context) {
	router := gin.Default()
	router.POST("/signup",controllers.SignUp)
	}

func loginRouter(c *gin.Context) {
	router := gin.Default()
	router.POST("/login",controllers.Login)

	router.Run(":8080")
}