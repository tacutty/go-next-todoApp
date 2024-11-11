package main

import (
	"go_next_todo/application/service"
	"go_next_todo/handler"
	repository "go_next_todo/infrastrucuture"
	"go_next_todo/infrastrucuture/db"
	router "go_next_todo/infrastrucuture/web"
	"go_next_todo/validator"
)

func main() {
	database := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(database.ConnectDB())
	userUsecase := service.NewUserUsecase(userRepository, userValidator)
	userHandler := handler.NewUserHandler(userUsecase)
	e := router.NewRouter(userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
