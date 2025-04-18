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
	envErr := godotenv.Load(".ENV")
	if envErr != nil {
		log.Println("error not read .env")
	}

	// Log JWT Secret for debugging
	jwtSecret := os.Getenv("JWT_SECRET")
	log.Printf("Loaded JWT_SECRET (length: %d): %s", len(jwtSecret), jwtSecret)

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
