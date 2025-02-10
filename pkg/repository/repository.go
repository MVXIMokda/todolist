package repository

import (
	"github.com/MVXIMokda/todolist"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GetUser(username, password string) (todo_app.User, error)
}

type ToDoList interface {
	Create(userId int, list todo_app.Todolist) (int, error)
	GetAll(userId int) ([]todo_app.Todolist, error)
	GetById(userId, listId int) (todo_app.Todolist, error)
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      NewTodoListPostgres(db),
	}
}
