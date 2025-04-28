package main

import (
	"Project-Golang/config"
	"Project-Golang/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func globalCorsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	envErr := godotenv.Load(".ENV")
	if envErr != nil {
		log.Println("error not read .env")
	}

	// Log JWT Secret for debugging
	jwtSecret := os.Getenv("JWT_SECRET")
	log.Printf("Loaded JWT_SECRET (length: %d)", len(jwtSecret))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Starting server...")
	config.ConnectDatabase()
	routes.SetupRoutes()

	log.Printf("Server http://localhost:%s 🚀", port)

	envErr = http.ListenAndServe(":"+port, globalCorsHandler(http.DefaultServeMux))
	if envErr != nil {
		log.Fatal("error server: ", envErr)
	}
}
