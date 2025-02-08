package service

import (
	todo_app "github.com/MVXIMokda/todolist"
	repository "github.com/MVXIMokda/todolist/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
