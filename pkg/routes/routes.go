package routes

import (
	"todo-app/pkg/controllers"
	"todo-app/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authorized"})
		})
		authorized.POST("/tasks", controllers.CreateTask)
		authorized.GET("/tasks", controllers.GetTasksByUserId)
		authorized.PUT("/tasks/:id", controllers.UpdateTask)
		authorized.DELETE("/tasks/:id", controllers.DeleteTask)
		authorized.GET("/tasks/export", controllers.ExportTasksExcel)
		authorized.GET("/tasks/:id", controllers.GetTaskById)
	}
}
