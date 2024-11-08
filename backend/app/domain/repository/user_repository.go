package repository

import (
	"go_next_todo/domain/model"
)

// IUserRepository interface
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
