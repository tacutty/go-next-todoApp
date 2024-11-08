package usecase

import (
	"go_next_todo/domain/model"
)

// IUserUsecase interface
type IUserUsecase interface {
	SignUp(user model.User) (model.User, error)
	Login(user model.User) (string, error)
}
