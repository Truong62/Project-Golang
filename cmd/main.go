package main

import (
	"Project-Golang/config"
	"Project-Golang/routes"
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
	log.Println("Starting server...")
	config.ConnectDatabase()
	routes.SetupRoutes()

	log.Printf("Server http://localhost:%s 🚀", port)
	envErr = http.ListenAndServe(":"+port, nil)
	if envErr != nil {
		log.Fatal("error server: ", envErr)
	}
}
