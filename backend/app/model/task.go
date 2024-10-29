package model

import "time"

type Task struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TaskResponse struct {
	ID          string `json:"id" gorm:"primary_key"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
