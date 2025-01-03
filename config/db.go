package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/ipaymupreviews/golang-gin-poc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// // Pastikan semua environment variables sudah ada
	// if host == "" || port == "" || user == "" || password == "" || dbname == "" {
	// 	return nil, fmt.Errorf("missing required environment variables")
	// }

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrasi tabel User
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate user table: %v", err)
	}

	// Migrasi tabel Otp
	if err := db.AutoMigrate(&models.Otp{}); err != nil {
		log.Fatalf("failed to migrate otp table: %v", err)
	}

	DB = db

}
