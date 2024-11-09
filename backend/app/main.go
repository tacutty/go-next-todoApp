package main

import (
	"fmt"
	"go_next_todo/application/service"
	"go_next_todo/handler"
	repository "go_next_todo/infrastrucuture"
	"go_next_todo/infrastrucuture/db"
	router "go_next_todo/infrastrucuture/web"
)

func main() {
	database := db.NewDB()
	userRepository := repository.NewUserRepository(database.ConnectDB())
	userUsecase := service.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	fmt.Println("check", userHandler)
	e := router.NewRouter(userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
