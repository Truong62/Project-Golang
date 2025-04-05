package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
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
}

func createDatabaseIfNotExists() {
	// Connect Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL server: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	dbName := os.Getenv("DB_NAME")

	// Create database if not found
	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Printf("Warning (ignore if already exists): %v", err)
	} else {
		log.Println("Database created:", dbName)
	}
}
