package models

import (
	"time"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primarykey"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	TaskName    string     `json:"task_name"`
	TaskDetails string     `json:"task_details"`
	Status      string     `json:"status"`
}

type UpdateTask struct {
	TaskName       string `json:"task_name"`
	TaskDetails    string `json:"task_details"`
	CompletionDate string `json:"completion_date"`
}

type UpdateStatus struct {
	Status string `json:"status"`
}
