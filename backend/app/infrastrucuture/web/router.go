package web

import (
	"net/http"
	"os"

	"go_next_todo/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uh handler.IUserHandler, th handler.ITaskHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge: 60,
	}))
	e.POST("/signup", uh.SignUp)
	e.POST("/login", uh.Login)
	e.POST("/logout", uh.Logout)
	e.GET("/csrf", uh.CsrfToken)
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", th.GetAllTasks)
	t.GET("/:taskId", th.GetTaskById)
	t.POST("", th.CreateTask)
	t.PUT("/:taskId", th.UpdateTask)
	t.DELETE("/:taskId", th.DeleteTask)
	return e
}
