package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type Users struct {
	Id uint `"gorm:primarykey"`
	Name string `"gorm:not null"`
	Email string `"gorm:unique;not null"`
	Password string `"gorm:not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


func (u *Users) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
}


func FindEmailUser (db *gorm.DB,email string) (*Users,error) {
	var user Users
	if err := db.Where("email=?",email).First(&user).Error; err != nil {
		return nil,err
	}
	return &user,nil
}

func CreateUser(db *gorm.DB,user *Users) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return errors.New("ハッシュ化失敗")
	}
	user.Password = string(hashed)

	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("ユーザー作成エラー:%w",err)
	}
	return nil
}
