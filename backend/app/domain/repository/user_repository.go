package repository

import (
	"go_next_todo/domain/model"
)

// IUserRepository interface
type IUserRepository interface {
	GetUserByNameAndEmail(user *model.User, username, email string) error
	CreateUser(user *model.User) error
}
