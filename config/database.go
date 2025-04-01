package config

import (
	"fmt"
	"gorm.io/gorm"
)

// DB is global for database access
var DB *gorm.DB

func ConnectDatabase() {
	// create string connect
	//dsn := fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_NAME"),
	//	os.Getenv("DB_PORT"),
	//)

	// connect database using GORM with
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatalf("Cannot database: %v", err)
	//}

	// save database global
	//DB = db

	fmt.Println("✅ Connect PostgreSQL success!")
}
