package todo_app

import "errors"

type Todolist struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"tittle" db:"tittle" binding:"required"`
	Description string `json:"description" db:"description"`
}

type Userslist struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int `json:"id"`
	ListId int `json:"list_id"`
	ItemId int `json:"item_id"`
}

type UpdateListInput struct {
	Title       *string `json:"tittle"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("Update structure has no values")
	}
	return nil
}
