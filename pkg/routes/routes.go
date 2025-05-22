package routes

import (
	"todo-app/pkg/controllers"
	"todo-app/pkg/middleware"
	"todo-app/pkg/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(cfg *models.Config) {
	auth := cfg.Gin.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp(cfg.Db))
		auth.POST("/signin", controllers.SignIn(cfg.Db))
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
			tasks.POST("", controllers.CreateTask(cfg.Db))
			tasks.GET("/export", controllers.ExportTasksExcel(cfg.Db))
			tasks.GET("/:id", controllers.GetTaskById(cfg.Db))
			tasks.PUT("/:id", controllers.UpdateTask(cfg.Db))
			tasks.DELETE("/:id", controllers.DeleteTask(cfg.Db))
		}

		users := authorized.Group("/users")
		{
			users.GET("/:id/tasks", controllers.GetTasksByUserId(cfg.Db))
		}
	}
}
