package infrastructure

import (
	"go_next_todo/domain/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// taskRepository struct
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository function
func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// GetAllTasks function
// Get all tasks
// @param tasks *[]model.Task
// @param userID string
// @return error
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userID string) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userID).Find(tasks).Error; err != nil {
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
func (tr *taskRepository) GetTaskByID(task *model.Task, taskID string, userID string) error {
	if err := tr.db.Joins("User").Where(("id = ? AND user_id = ?"), taskID, userID).First(task).Error; err != nil {
		return err
	}
	return nil
}

// CreateTask function
// Create task
// @param task *model.Task
// @return error
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTask function
// Update task
// @param task *model.Task
// @param taskID string
// @param userID string
// @return error
func (tr *taskRepository) UpdateTask(task *model.Task, taskID string, userID string) error {
	if err := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", taskID, userID).Updates(map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"completed":   task.Completed,
	}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteTask function
// Delete task
// @param taskID string
// @param userID string
// @return error
func (tr *taskRepository) DeleteTask(taskID string, userID string) error {
	if err := tr.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}
