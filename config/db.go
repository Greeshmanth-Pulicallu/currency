package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	dsn := os.Getenv("DSN")

	if dsn == "" {
		log.Println("invalid dsn")
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
}
