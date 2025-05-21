package env

import (
	"todo-app/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	Db  *gorm.DB
	Gin *gin.Engine
}

func GetConfig() *Config {
	models.Connect()
	models.InitModels()

	return &Config{
		Db:  models.GetDB(),
		Gin: gin.Default(),
	}
}
