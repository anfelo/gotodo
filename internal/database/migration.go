package database

import (
	"github.com/anfelo/gotodo/internal/todos"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&todos.TodoList{}, &todos.Todo{}); err != nil {
		return err
	}
	return nil
}
