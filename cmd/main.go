package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error not read .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!!! Go running ... 🚀")
	})

	log.Printf("Server http://localhost:%s 🚀", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("error server: ", err)
	}
}
