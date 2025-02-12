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
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type ToDoItem interface {
	Create(userId int, listId todo_app.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo_app.TodoItem, error)
	GetById(userId, itemId int) (todo_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_app.UpdateItemInput) error
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
		ToDoItem:      NewTodoItemPostgres(db),
	}
}
