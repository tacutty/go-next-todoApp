package service

import (
	"go_next_todo/application/usecase"
	"go_next_todo/domain/model"
	"go_next_todo/domain/repository"
	"go_next_todo/utils"
	"go_next_todo/validator"
)

// taskUsecase struct
type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

// NewTaskUsecase function
func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) usecase.ITaskUsecase {
	return &taskUsecase{tr, tv}
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
	validatorErr := tu.tv.TaskValidate(task)
	if validatorErr != nil {
		return model.TaskResponse{}, validatorErr
	}

	id, err := utils.GeneULIDString()
	if err != nil {
		return model.TaskResponse{}, err
	}
	task.ID = id

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
	validatorErr := tu.tv.TaskValidate(task)
	if validatorErr != nil {
		return model.TaskResponse{}, validatorErr
	}

	if err := tu.tr.UpdateTask(&task, taskId, userId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:          taskId,
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
