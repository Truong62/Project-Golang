package repository

import (
	"example.com/crud-go/config"
	"example.com/crud-go/models"
)

type TodoRepository struct{}

func (r *TodoRepository) Create(todo *models.Todo) error {
	result := config.DB.Create(todo)
	return result.Error
}
