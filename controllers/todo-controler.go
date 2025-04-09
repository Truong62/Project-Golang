package controllers

import (
	"Project-Golang/models"
	"Project-Golang/repository"
	"encoding/json"
	"fmt"
	"log"
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

	log.Printf("DEBUG: Decoded todo: %+v", todo)

	if err := todo.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Invalid data: %s", err.Error()), http.StatusBadRequest)
		return
	}

	isErr := todoRepo.Create(&todo)

	if isErr != nil {
		log.Printf("DEBUG: Error creating todo: %v", isErr)
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
