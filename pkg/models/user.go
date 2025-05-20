package models

import (
	"errors"
	"todo-app/pkg/config"
	"todo-app/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Tasks    []Task `json:"tasks,omitempty"`
}

// var db *gorm.DB

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(&User{}, &Task{})
// }

func (u *User) SignUp() (*User, error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword
	db := config.GetDB()
	result := db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (u *User) SignIn(password string) (string, error) {
	db := config.GetDB()
	if err := db.Where("username = ?", u.Username).First(u).Error; err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, u.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
