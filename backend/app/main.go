package main

import (
	"go_next_todo/db"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	db := db.NewDB()
	conn := db.ConnectDB()
	if conn == nil {
		log.Fatalln("Failed to connect to database: connection is nil")
	}
	log.Println("Database connection established:", conn)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
