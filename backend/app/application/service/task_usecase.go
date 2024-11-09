package service

import (
	"go_next_todo/application/usecase"
	"go_next_todo/domain/model"
	"go_next_todo/domain/repository"
)

// taskUsecase struct
type taskUsecase struct {
	tr repository.ITaskRepository
}

// NewTaskUsecase function
func NewTaskUsecase(tr repository.ITaskRepository) usecase.ITaskUsecase {
	return &taskUsecase{tr}
}

// GetAllTasks function
// @param userId string
// @return []model.TaskResponse, error
func (tu *taskUsecase) GetAllTasks(userId string) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Completed:   v.Completed,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

// GetTaskById function
// @param userId string
// @param taskId string
// @return model.TaskResponse, error
func (tu *taskUsecase) GetTaskById(userId, taskId string) (model.TaskResponse, error) {
	var task model.Task
	if err := tu.tr.GetTaskByID(&task, taskId, userId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
	return resTask, nil
}

// CreateTask function
// @param task model.Task
// @return model.TaskResponse, error
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
	return resTask, nil
}

// UpdateTask function
// @param task model.Task
// @param userId string
// @param taskId string
// @return model.TaskResponse, error
func (tu *taskUsecase) UpdateTask(task model.Task, userId, taskId string) (model.TaskResponse, error) {
	if err := tu.tr.UpdateTask(&task, taskId, userId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
	return resTask, nil
}

// DeleteTask function
// @param userId string
// @param taskId string
// @return error
func (tu *taskUsecase) DeleteTask(userId, taskId string) error {
	if err := tu.tr.DeleteTask(taskId, userId); err != nil {
		return err
	}
	return nil
}
