package controllers

import (
	"net/http"
	"todo-app/pkg/models"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user := &models.User{
		Username: input.Username,
		Password: input.Password,
	}

	createdUser, err := user.SignUp()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       createdUser.ID,
		"username": createdUser.Username,
	})
}

func SignIn(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user := &models.User{
		Username: input.Username,
	}

	token, err := user.SignIn(input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
