package domain

type Repository interface {
	CreateTodo(todo *Todo) error
	GetTodos(title string, isDone *bool) ([]*Todo, error)
	GetTodo(id uint) (*Todo, error)
	UpdateTodo(condition *Todo, instance *Todo) error
	DeleteTodo(id uint) error
}
