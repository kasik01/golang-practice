package routes

import (
	"todo-app/pkg/controllers"
	"todo-app/pkg/env"
	"todo-app/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(cfg *env.Config) {
	auth := cfg.Gin.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/signin", controllers.SignIn)
		auth.POST("/renew-token", controllers.RefreshToken)
	}

	authorized := cfg.Gin.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authorized"})
		})

		tasks := authorized.Group("/tasks")
		{
			tasks.POST("", controllers.CreateTask)
			tasks.GET("/export", controllers.ExportTasksExcel)
			tasks.GET("/:id", controllers.GetTaskById)
			tasks.PUT("/:id", controllers.UpdateTask)
			tasks.DELETE("/:id", controllers.DeleteTask)
		}

		users := authorized.Group("/users")
		{
			users.GET("/:id/tasks", controllers.GetTasksByUserId)
		}
	}

}
