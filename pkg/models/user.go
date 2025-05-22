package models

import (
	"errors"
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

func (u *User) SignUp(db *gorm.DB) (*User, error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword
	result := db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (u *User) SignIn(password string, db *gorm.DB) (*utils.TokenResponse, error) {
	if err := db.Where("username = ?", u.Username).First(u).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, u.Password) {
		return nil, errors.New("invalid password")
	}

	return utils.GenerateToken(u.ID)
}
