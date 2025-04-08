package repository

import (
	"Project-Golang/config"
	"Project-Golang/models"
)

type TodoRepository struct{}

func (r *TodoRepository) Create(todo *models.Todo) error {
	result := config.DB.Create(todo)
	return result.Error
}
