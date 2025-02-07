package todo_app

type Todolist struct {
	Id          int    `json:"id"`
	Tittle      string `json:"title"`
	Description string `json:"description"`
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
