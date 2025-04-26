package routes

import (
	"Project-Golang/controllers"
	"Project-Golang/middleware"
	"fmt"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		handler := middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				controllers.CreateTodoList(w, r)
			case http.MethodGet:
				controllers.GetTodoList(w, r)
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})
		handler(w, r)
	})

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/todos" || path == "/todos/" {
			return
		}

		handler := middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPut:
				controllers.PutTodoById(w, r)
			case http.MethodDelete:
				controllers.DeleteTodoById(w, r)
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})
		handler(w, r)
	})

	// Public route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		_, err := fmt.Fprintln(w, "Hello!!! Go running ... 🚀")
		if err != nil {
			return
		}
	})
}
