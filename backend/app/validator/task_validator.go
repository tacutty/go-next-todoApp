package validator

import (
	"go_next_todo/domain/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("Title is required"),
			validation.RuneLength(1, 10).Error("Title must be between 1 and 30 characters"),
		),
		validation.Field(
			&task.Description,
			validation.Required.Error("Description is required"),
			validation.RuneLength(1, 100).Error("Description must be between 1 and 100 characters"),
		),
		validation.Field(
			&task.Completed,
			validation.In(true, false).Error("Completed must be true or false"),
		),
	)
}
