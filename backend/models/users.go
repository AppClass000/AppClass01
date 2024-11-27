package models

import (
	
	
	"gorm.io/driver/mysql"
)


type Users struct {
	id uint "gorm:primarykey"
	name string 
	email string
	password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func users() {
	dsn := "root:sk3316624@tcp(127.0.0.1:3306)/appclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
		log.fatal("データベース接続error:",err)
	}

	newUser := Users{name:"kaito"}
	
}