package routes

import (
	"Project-Golang/controllers"
	"fmt"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateTodoList(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello!!! Go running ... 🚀")
		if err != nil {
			return
		}
	})
}
