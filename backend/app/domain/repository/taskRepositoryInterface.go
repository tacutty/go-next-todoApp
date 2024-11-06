package repository

import (
	"go_next_todo/domain/model"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userID string) error
	GetTaskByID(task *model.Task, taskID string, userID string) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, taskID string, userID string) error
	DeleteTask(taskID string, userID string) error
}