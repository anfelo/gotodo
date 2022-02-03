package todos

import "gorm.io/gorm"

// Service - the struct for our todos service
type Service struct {
	DB *gorm.DB
}

// Todo - defines the Todo model
type Todo struct {
	gorm.Model
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
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetTodo - retrieves a Todo by their ID from the db
func (s *Service) GetTodo(ID uint) (Todo, error) {
	var todo Todo
	if result := s.DB.First(&todo, ID); result.Error != nil {
		return Todo{}, result.Error
	}
	return todo, nil
}

// CreateTodo - adds a new Todo to the database
func (s *Service) CreateTodo(todo Todo) (Todo, error) {
	if result := s.DB.Save(&todo); result.Error != nil {
		return Todo{}, result.Error
	}
	return todo, nil
}

// UpdateTodo - updates a Todo by ID with new Todo info
func (s *Service) UpdateTodo(ID uint, newTodo Todo) (Todo, error) {
	todo, err := s.GetTodo(ID)
	if err != nil {
		return Todo{}, err
	}

	if result := s.DB.Model(&todo).Updates(newTodo); result.Error != nil {
		return Todo{}, result.Error
	}

	return todo, nil
}

// DeleteTodo - deletes a Todo from the database by ID
func (s *Service) DeleteTodo(ID uint) error {
	if result := s.DB.Delete(&Todo{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllTodos - retrieves all Todos from the db
func (s *Service) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	if result := s.DB.Find(&todos); result.Error != nil {
		return []Todo{}, result.Error
	}
	return todos, nil
}
