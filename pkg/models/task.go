package models

import (
	"time"
	"todo-app/pkg/config"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `gorm:"default:pending" json:"status"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
}

func (t *Task) CreateTask(userId uint) (*Task, error) {
	t.UserID = userId
	db := config.GetDB()
	result := db.Create(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func (t *Task) GetTaskByUserId(userId uint) ([]Task, error) {
	db := config.GetDB()
	var tasks []Task
	result := db.Where("user_id = ?", userId).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (t *Task) GetTaskById(taskId uint) (*Task, error) {
	db := config.GetDB()
	var task Task
	result := db.First(&task, taskId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (t *Task) UpdateTask(taskId uint) (*Task, error) {
	db := config.GetDB()
	var task Task
	result := db.First(&task, taskId)
	if result.Error != nil {
		return nil, result.Error
	}

	task.Title = t.Title
	task.Description = t.Description
	task.DueDate = t.DueDate
	task.Status = t.Status

	result = db.Save(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (t *Task) DeleteTask(taskId uint) error {
	db := config.GetDB()
	result := db.Delete(&Task{}, taskId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
