package main

import (
	todo_app "github.com/MVXIMokda/todolist"
	"github.com/MVXIMokda/todolist/pkg/handler"
	"github.com/MVXIMokda/todolist/pkg/repository"
	"github.com/MVXIMokda/todolist/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	r := repository.NewRepository()
	services := service.NewService(r)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
