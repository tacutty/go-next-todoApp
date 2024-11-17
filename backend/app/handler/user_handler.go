package handler

import (
	"fmt"
	"go_next_todo/application/usecase"
	"go_next_todo/domain/model"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
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

	token, err := uh.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// Logout function
// Logout user
// @param c echo.Context
// @return error
func (uh *userHandler) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// CsrfToken function
// Get csrf token
// @param c echo.Context
// @return error
func (uh *userHandler) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf": token,
	})
}
