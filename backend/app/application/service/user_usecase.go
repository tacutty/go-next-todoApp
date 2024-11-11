package service

import (
	"go_next_todo/application/usecase"
	"go_next_todo/domain/model"
	"go_next_todo/domain/repository"
	"go_next_todo/utils"
	"go_next_todo/validator"

	"golang.org/x/crypto/bcrypt"
)

// userUsecase struct
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// NewUserUsecase function
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) usecase.IUserUsecase {
	return &userUsecase{ur, uv}
}

// SignUp function
// Sign up
// @param user model.User
// @return model.User, error
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	validatorErr := uu.uv.UserValidate(user)
	if validatorErr != nil {
		return model.UserResponse{}, validatorErr
	}
	id, err := utils.GeneULIDString()
	if err != nil {
		return model.UserResponse{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
	}

	return resUser, nil
}

// Login function
// Login
// @param user model.User
// @return string, error
func (uu *userUsecase) Login(user model.User) (string, error) {
	validatorErr := uu.uv.UserValidate(user)
	if validatorErr != nil {
		return "", validatorErr
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByNameAndEmail(&storedUser, user.Username, user.Email); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	token, err := utils.GenerateJwtToken(storedUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
