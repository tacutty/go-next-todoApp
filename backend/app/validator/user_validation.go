package validator

import (
	"go_next_todo/domain/model"

	"github.com/go-ozzo/ozzo-validation/v4/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// IUserValidator interface
type IUserValidator interface {
	UserValidate(user model.User) error
}

// userValidator struct
type userValidator struct{}

// NewUserValidator function
func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// UserValidate function
func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Username,
			validation.Required.Error("Username is required"),
			validation.RuneLength(4, 10).Error("Username must be between 4 and 10 characters"),
		),
		validation.Field(&user.Email,
			validation.Required.Error("Email is required"),
			validation.RuneLength(1, 30).Error("Email must be between 4 and 10 characters"),
			is.Email.Error("Email is not valid"),
		),
		validation.Field(&user.Password,
			validation.Required.Error("Password is required"),
			validation.RuneLength(6, 10).Error("Password must be between 6 and 10 characters"),
		),
	)
}
