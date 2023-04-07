package domain

type Service interface {
	AddTodo(title, content string, isDone *bool) error
	ListTodo(title string, isDone *bool) ([]*Todo, error)
	GetTodo(id uint) (*Todo, error)
	UpdateTodo(id uint, title, content string, isDone *bool) error
	DeleteTodo(id uint) error
}
