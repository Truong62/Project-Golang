package repository

import (
	"Project-Golang/config"
	"Project-Golang/models"
	"strings"
)

type TodoRepository struct{}

func (r *TodoRepository) Create(todo *models.Todo) error {
	result := config.DB.Create(todo)
	return result.Error
}

func (r *TodoRepository) GetAllTodos(todos *[]models.Todo) error {
	result := config.DB.Find(todos)
	return result.Error
}

func (r *TodoRepository) GetTodosBySearch(keySearch string, todos *[]models.Todo) error {
	keySearch = strings.TrimSpace(keySearch)
	query := config.DB.Where("title LIKE ?", "%"+keySearch+"%")

	result := query.Find(todos)

	var allTodos []models.Todo
	config.DB.Find(&allTodos)
	return result.Error
}
