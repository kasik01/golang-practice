package main

import (
	"fmt"
	"todo-app/pkg/config"
	"todo-app/pkg/models"
	"todo-app/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.Connect()
	models.InitModels()
	router := gin.Default()

	routes.RegisterRoutes(router)

	port := config.GetAppConfig().APP_PORT
	fmt.Println("Server running on port:", port)
	router.Run(":" + port)
}
