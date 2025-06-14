package models

import (
	"time"

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
	db := GetDB()
	t.UserID = userId
	result := db.Create(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func (t *Task) GetTaskByUserId(userId uint) ([]Task, error) {
	db := GetDB()
	var tasks []Task
	result := db.Where("user_id = ?", userId).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (t *Task) GetTaskById(taskId uint) (*Task, error) {
	db := GetDB()
	var task Task
	result := db.First(&task, taskId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (t *Task) UpdateTask(db *gorm.DB, taskId uint) (*Task, error) {
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
	db := GetDB()
	result := db.Delete(&Task{}, taskId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
