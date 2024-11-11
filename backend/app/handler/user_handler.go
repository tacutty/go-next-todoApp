package handler

import (
	"fmt"
	"go_next_todo/application/usecase"
	"go_next_todo/domain/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

// UserHandler struct
type userHandler struct {
	uu usecase.IUserUsecase
}

// NewUserHandler function
func NewUserHandler(uu usecase.IUserUsecase) IUserHandler {
	return &userHandler{uu}
}

// CreateUser function
// Create user
// @param c echo.Context
// @return error
func (uh *userHandler) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uh.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

// Login function
// Login user
// @param c echo.Context
// @return error
func (uh *userHandler) Login(c echo.Context) error {
	fmt.Println("login")
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("user", user)
	token, err := uh.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
