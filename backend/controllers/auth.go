package controllers

import (
	"net/http"
	"time"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

 

type Users struct {
	Id uint `"gorm:primarykey" "json:id"`
	Name string  `"gorm:not null" "json:name" "binding:required"`
	Email string `"gorm:unique; not null" "json:email" "binding:required,email"`
	Password string `"gorm:not null" "json:password" "binding:required"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}

var db *gorm.DB
func InitDB () {
	dsn := "root:sk3316624@tcp(127.0.0.1:3306)/appclass?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
    db,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベース接続エラー",err)
	}

	err = db.AutoMigrate(&Users{})
	if err != nil {
	   	log.Fatalf("MigrationError",err)
	   }
}


func Register(c *gin.Context) {
	var input Users

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		log.Fatalf("パスワード暗号化error",err)
	}

	newUser :=Users{
		Name: input.Name,
		Email: input.Email,
		Password: hashedPassword,
        }

	if err := db.Create(&newUser).Error ; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"データベースに登録できませんでした"})
	return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":"登録されました",
		"user":newUser,
	})
}


func hashPassword(password string) (string,error)  {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}
	return string(hashed),nil
}



