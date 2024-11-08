package usecase

import (
	"go_next_todo/domain/model"
)

type ITaskRepository interface {
	CreateTask(task *model.Task) error
	GetAllTasks(tasks *[]model.Task, userID string) error
	GetTaskByID(task *model.Task, taskID string, userID string) error
	UpdateTask(task *model.Task, taskID string, userID string) error
	DeleteTask(taskID string, userID string) error
}