package models

import (
	"todo-app/pkg/config"
)

func InitModels() {
	db := config.GetDB()
	db.AutoMigrate(&User{}, &Task{})
}
