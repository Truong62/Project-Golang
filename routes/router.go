package routes

import (
	"Project-Golang/controllers"
	"fmt"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.CreateTodoList(w, r)
		case http.MethodGet:
			controllers.GetTodoList(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			controllers.PutTodoById(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello!!! Go running ... 🚀")
		if err != nil {
			return
		}
	})
}
