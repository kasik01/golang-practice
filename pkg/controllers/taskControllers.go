package controllers

import (
	"net/http"
	"strconv"

	"todo-app/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	createdTask, err := task.CreateTask(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          createdTask.ID,
		"title":       createdTask.Title,
		"description": createdTask.Description,
		"due_date":    createdTask.DueDate,
		"status":      createdTask.Status,
		"user_id":     createdTask.UserID,
	})
}

func GetTasksByUserId(c *gin.Context) {
	// userIDVal, exists := c.Get("user_id") // get id from token

	userIDParam := c.Param("id")
	if userIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userID := uint(userIDUint64)

	var task models.Task
	tasks, err := task.GetTaskByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	taskIDParam := c.Param("id")
	if taskIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}
	taskIDUint64, err := strconv.ParseUint(taskIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	taskID := uint(taskIDUint64)
	var task models.Task
	taskDetails, err := task.GetTaskById(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskDetails)
}

func UpdateTask(c *gin.Context) {
	taskIDParam := c.Param("id")
	if taskIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	taskIDUint64, err := strconv.ParseUint(taskIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	taskID := uint(taskIDUint64)

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := task.UpdateTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	taskIDParam := c.Param("id")
	if taskIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}
	taskIDUint64, err := strconv.ParseUint(taskIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	taskID := uint(taskIDUint64)

	var task models.Task
	err = task.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func ExportTasksExcel(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	var task models.Task
	tasks, err := task.GetTaskByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	headers := []string{"Tiêu đề", "Mô tả", "Ngày kết thúc", "Trạng thái"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1) // row 1
		f.SetCellValue(sheet, cell, h)
	}

	// Content
	for i, task := range tasks {
		row := i + 2 // bắt đầu từ dòng 2

		f.SetCellValue(sheet, "A"+strconv.Itoa(row), task.Title)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), task.Description)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), task.DueDate.Format("2006-01-02 15:04:05"))
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), task.Status)
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", `attachment; filename="danh sach cv.xlsx"`)
	c.Header("Content-Transfer-Encoding", "binary")

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate excel"})
		return
	}

}
