package main

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
	
)



func main() {
	controllers.InitDB()

	router := gin.Default()
	router.POST("/signup",controllers.Register)
	router.Run(":8080")
}