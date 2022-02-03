package todos

// Service - the struct for our todos service
type Service struct{}

// Todo - defines the Todo model
type Todo struct {
	Description string
}

// TodoService - the interface for our Todo service
type TodoService interface {
	GetTodo(ID uint) (Todo, error)
	CreateTodo(Todo Todo) (Todo, error)
	UpdateTodo(ID uint, newTodo Todo) (Todo, error)
	DeleteTodo(ID uint) error
	GetAllTodos() ([]Todo, error)
}

// NewService - returns new todos service
func NewService() *Service {
	return &Service{}
}

// GetAllTodos - retrieves all Todos from the db
func (s *Service) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	return todos, nil
}
