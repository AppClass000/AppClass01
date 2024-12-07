package main

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
	
	
)



func main() {
	controllers.InitDB()
	router := gin.Default()


	router.Static("/public","/frontend/src/component")
	router.StaticFile("/","/frontend/public/index.html")
    router.POST("/login",controllers.Login)
	router.POST("/signup",controllers.SignUp)


	router.Run(":8080")

}
