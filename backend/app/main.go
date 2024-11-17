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
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(database.ConnectDB())
	userUsecase := service.NewUserUsecase(userRepository, userValidator)
	taskUsecase := service.NewTaskUsecase(taskRepository, taskValidator)
	userHandler := handler.NewUserHandler(userUsecase)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := router.NewRouter(userHandler, taskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
