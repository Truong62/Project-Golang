package controllers

import (
	"Project-Golang/models"
	"Project-Golang/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	log.Printf("DEBUG: Successfully created todo with ID: %d", todo.ID)
}

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	keySearch := r.URL.Query().Get("search")
	log.Printf("DEBUG: Search parameter: %s", keySearch)

	w.Header().Set("Content-Type", "application/json")

	if keySearch != "" {
		if err := todoRepo.GetTodosBySearch(keySearch, &todos); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
	} else {
		if err := todoRepo.GetAllTodos(&todos); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, "Error encode data", http.StatusInternalServerError)
		return
	}
}

func PutTodoById(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id := parts[len(parts)-1]

	log.Printf("DEBUG: Updating todo with ID: %s", id)

	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := todo.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Invalid data: %s", err.Error()), http.StatusBadRequest)
		return
	}

	isErr := todoRepo.UpdateTodoById(id, &todo)

	if isErr != nil {
		http.Error(w, "Error updating todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data:    todo,
	}); err != nil {
		return
	}
}

func DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id := parts[len(parts)-1]

	err := todoRepo.DeleteTodoById(id)
	if err != nil {
		http.Error(w, "Error deleting todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data:    "Todo deleted successfully",
	}); err != nil {
		return
	}
}
