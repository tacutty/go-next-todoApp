package service

import (
	"go_next_todo/domain/model"
	"go_next_todo/domain/repository"
	"go_next_todo/application/usecase"
)

// userUsecase struct
type userUsecase struct {
	ur repository.IUserRepository
}

// NewUserUsecase function
func NewUserUsecase(ur repository.IUserRepository) usecase.IUserUsecase {
	return &userUsecase{ur}
}

// SignUp function
// Sign up
// @param user model.User
// @return model.User, error
func (uu *userUsecase) SignUp(user model.User) (model.User, error) {
	if err := uu.ur.CreateUser(&user); err != nil {
		return user, err
	}
	return user, nil
}

// Login function
// Login
// @param user model.User
// @return string, error
func (uu *userUsecase) Login(user model.User) (string, error) {
	return "", nil
}
