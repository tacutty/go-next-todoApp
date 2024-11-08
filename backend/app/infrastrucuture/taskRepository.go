package infrastructure

import (
	"go_next_todo/domain/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) GetAllTasks(tasks *[]model.Task, userID string) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userID).Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) GetTaskByID(task *model.Task, taskID string, userID string) error {
	if err := tr.db.Joins("User").Where(("id = ? AND user_id = ?"), taskID, userID).First(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) UpdateTask(task *model.Task, taskID string, userID string) error {
	if err := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", taskID, userID).Updates(map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"completed":   task.Completed,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteTask(taskID string, userID string) error {
	if err := tr.db.Where ("id = ? AND user_id = ?", taskID, userID).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}
