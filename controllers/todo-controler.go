package controllers

import (
	"encoding/json"
	"example.com/crud-go/models"
	"example.com/crud-go/repository"
	"fmt"
	"net/http"
)

var todoRepo = &repository.TodoRepository{}

func CreateTodoList(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	// Decode JSON to struct
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := todo.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Invalid data: %s", err.Error()), http.StatusBadRequest)
		return
	}

	isErr := todoRepo.Create(&todo)

	if isErr != nil {
		http.Error(w, "error create tode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(todo)
	if err != nil {
		return
	}
}
