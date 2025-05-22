package main

import (
	"fmt"

	"todo-app/pkg/config"
	"todo-app/pkg/models"
	"todo-app/pkg/routes"
)

func main() {
	config.LoadEnv()
	// models.Connect()
	// models.InitModels()
	// router := gin.Default()

	cfg := models.GetConfig()

	routes.RegisterRoutes(cfg)

	port := config.GetAppConfig().APP_PORT
	fmt.Println("Server running on port:", port)
	cfg.Gin.Run(":" + port)
}
