package controllers

import (
	"encoding/json"
	"example.com/crud-go/models"
	"example.com/crud-go/repository"
	"net/http"
)

var todoRepo = &repository.TodoRepository{}

func CreateTodoList(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := todoRepo.Create(&todo)

	if err != nil {
		http.Error(w, "error create tode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		return
	}
}
