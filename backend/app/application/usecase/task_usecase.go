package usecase

import (
	"go_next_todo/domain/model"
)

type ITaskUsecase interface {
	GetAllTasks(userId string) ([]model.TaskResponse, error)
	GetTaskById(userId string, taskId string) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId string, taskId string) (model.TaskResponse, error)
	DeleteTask(userId string, taskId string) error
}