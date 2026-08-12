package Database

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConDb() {
	dsn := "host=localhost user=gintest password=123456 dbname=mydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("do not connect to database", err)
	}
	log.Println("Connected")
	// err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("connection break", err)
	}
	DB = db
}
