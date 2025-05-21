package main

import (
	"fmt"

	"todo-app/pkg/config"
	"todo-app/pkg/env"
	"todo-app/pkg/routes"
)

func main() {
	config.LoadEnv()

	cfg := env.GetConfig()
	routes.RegisterRoutes(cfg)

	port := config.GetAppConfig().APP_PORT
	fmt.Println("Server running on port:", port)
	cfg.Gin.Run(":" + port)
}
