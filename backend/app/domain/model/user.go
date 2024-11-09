package model

import (
	"time"
)

// User struct
type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// UserResponse struct
type UserResponse struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"unique not null"`
}
