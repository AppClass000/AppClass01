package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp (c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}
	if err := models.CreateUser(db,&input); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":"ユーザー作成失敗"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"ユーザー作成に成功しました"})
}


func Login(c *gin.Context) {
	email := c.Query("email")
	if email == ""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"メールが見つかりません"})
		return
	}
	user,err := models.FindEmailUser(db,email)
	if err != nil {
		c.JSON(404,gin.H{"error":"ユーザーが見つかりません"})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{"message":"ログインに成功",
	            "user":user})
}