package Database

import (
	"fmt"
	"log"
	"os"
	"project/models"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConDb() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5433"),
		getEnv("DB_USER", "gintest"),
		getEnv("DB_PASSWORD", "123456"),
		getEnv("DB_NAME", "mybd"),
		getEnv("DB_SSLMODE", "disable"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("do not connect to database", err)
	}
	log.Println("Connected")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("migration break", err)
	}
	log.Println("migration Done")
	DB = db
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
