package web

import (
	"go_next_todo/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc handler.IUserHandler) *echo.Echo {
	e := echo.New()

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)

	return e
}
