package service

import (
	"go_next_todo/domain/model"
	"go_next_todo/domain/repository"
)

// taskUsecase struct
type taskUsecase struct {
	tr repository.ITaskRepository
}

// NewTaskUsecase function
func NewTaskUsecase(tr repository.ITaskRepository) *taskUsecase {
	return &taskUsecase{tr}
}

// CreateTask function
// Create task
// @param task model.Task
// @return model.Task, error
func (tu *taskUsecase) CreateTask(task model.Task) (model.Task, error) {
	if err := tu.tr.CreateTask(&task); err != nil {
		return task, err
	}
	return task, nil
}

// GetAllTasks function
// Get all tasks
// @param tasks *[]model.Task
// @param userID string
// @return error
func (tu *taskUsecase) GetAllTasks(tasks *[]model.Task, userID string) error {
	if err := tu.tr.GetAllTasks(tasks, userID); err != nil {
		return err
	}
	return nil
}

// GetTaskByID function
// Get task by ID
// @param task *model.Task
// @param taskID string
// @param userID string
// @return error
func (tu *taskUsecase) GetTaskByID(task *model.Task, taskID string, userID string) error {
	if err := tu.tr.GetTaskByID(task, taskID, userID); err != nil {
		return err
	}
	return nil
}

// UpdateTask function
// Update task
// @param task model.Task
// @param taskID string
// @param userID string
// @return error
func (tu *taskUsecase) UpdateTask(task model.Task, taskID string, userID string) error {
	if err := tu.tr.UpdateTask(&task, taskID, userID); err != nil {
		return err
	}
	return nil
}

// DeleteTask function
// Delete task
// @param taskID string
// @param userID string
// @return error
func (tu *taskUsecase) DeleteTask(taskID string, userID string) error {
	if err := tu.tr.DeleteTask(taskID, userID); err != nil {
		return err
	}
	return nil
}
