package config

import (
	"Project-Golang/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is global for database access
var DB *gorm.DB

func ConnectDatabase() {
	createDatabaseIfNotExists()

	//create string connect
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	//connect database using GORM with
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot database: %v", err)
	}

	//save database global
	DB = db

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		return
	}
}

func createDatabaseIfNotExists() {
	// Connect to Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
	)

	drivenName := "postgres"
	db, err := sql.Open(drivenName, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL server: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	dbName := os.Getenv("DB_NAME")

	// Check if the database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1)", dbName).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	if exists {
		log.Printf("Database already exists: %s", dbName)
		return
	}

	// Create database if not found
	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Printf("Error creating database: %v", err)
	} else {
		log.Println("Database created:", dbName)
	}
}
