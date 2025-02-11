package service

import (
	"github.com/MVXIMokda/todolist"
	"github.com/MVXIMokda/todolist/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, list todo_app.Todolist) (int, error)
	GetAll(userId int) ([]todo_app.Todolist, error)
	GetById(userId, listId int) (todo_app.Todolist, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type ToDoItem interface {
	Create(userId, listId int, item todo_app.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo_app.TodoItem, error)
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList:      NewTodoListService(repos.ToDoList),
		ToDoItem:      NewTodoItemService(repos.ToDoItem, repos.ToDoList),
	}
}
