package main

import (
	todo_app "github.com/MVXIMokda/todolist"
	"github.com/MVXIMokda/todolist/pkg/handler"
	"github.com/MVXIMokda/todolist/pkg/repository"
	"github.com/MVXIMokda/todolist/pkg/service"
	"log"
)

func main() {
	r := repository.NewRepository()
	services := service.NewService(r)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
