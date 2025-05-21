package routes

import (
	"todo-app/pkg/controllers"
	"todo-app/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/signin", controllers.SignIn)
		auth.POST("/renew-token", controllers.RefreshToken)
	}

	authorized := r.Group("/")
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
