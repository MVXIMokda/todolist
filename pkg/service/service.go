package service

import repository "github.com/MVXIMokda/todolist/pkg/repository"

type Authorization interface {
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

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
