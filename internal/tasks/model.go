package tasks

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Description string `json:"description"`
	Status      string `json:"status"`
	Epic        string `json:"epic"`
	Daily       bool   `json:"daily"`
}

func NewTask(description, status, epic string, daily bool) *Task {
	return &Task{
		Description: description,
		Status:      status,
		Epic:        epic,
		Daily:       daily,
	}
}
