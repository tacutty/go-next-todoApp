package handler

import (
	"go_next_todo/application/usecase"
	"go_next_todo/domain/model"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITaskHandler interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

// TaskHandler struct
type taskHandler struct {
	tu usecase.ITaskUsecase
}

// NewTaskHandler function
func NewTaskHandler(tu usecase.ITaskUsecase) ITaskHandler {
	return &taskHandler{tu}
}

// GetAllTasks function
// Get all tasks
// @param c echo.Context
// @return error
func (th *taskHandler) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	tasks, err := th.tu.GetAllTasks(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

// GetTaskById function
// Get task by id
// @param c echo.Context
// @return error
func (th *taskHandler) GetTaskById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	taskId := c.Param("taskId")

	task, err := th.tu.GetTaskById(userId, taskId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

// CreateTask function
// Create task
// @param c echo.Context
// @return error
func (th *taskHandler) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task.UserID = userId

	taskRes, err := th.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, taskRes)
}

// UpdateTask function
// Update task
// @param c echo.Context
// @return error
func (th *taskHandler) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	taskId := c.Param("taskId")

	taskRes, err := th.tu.UpdateTask(task, userId, taskId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

// DeleteTask function
// Delete task
// @param c echo.Context
// @return error
func (th *taskHandler) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	taskId := c.Param("taskId")

	if err := th.tu.DeleteTask(userId, taskId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
