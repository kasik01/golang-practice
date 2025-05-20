package main

import (
	"fmt"
	"todo-app/pkg/config"
	"todo-app/pkg/models"
	"todo-app/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	models.InitModels()

	router := gin.Default()

	routes.RegisterRoutes(router)

	fmt.Println("Server running on port 8080")
	router.Run(":8080")
}
