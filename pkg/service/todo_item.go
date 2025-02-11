package service

import (
	"github.com/MVXIMokda/todolist"
	"github.com/MVXIMokda/todolist/pkg/repository"
)

type TodoItemService struct {
	repo     repository.ToDoItem
	listRepo repository.ToDoList
}

func NewTodoItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo_app.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo_app.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
