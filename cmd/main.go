package main

import (
	"example.com/crud-go/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println("error not read .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	config.ConnectDatabase()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello!!! Go running ... 🚀")
		if err != nil {
			return
		}
	})

	log.Printf("Server http://localhost:%s 🚀", port)
	envErr = http.ListenAndServe(":"+port, nil)
	if envErr != nil {
		log.Fatal("error server: ", envErr)
	}
}
