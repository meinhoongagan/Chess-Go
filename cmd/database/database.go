package database

import (
	"Chess-Go/config"
	"Chess-Go/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.AppConfig.Database.URL
	if dsn == "" {
		log.Fatal("❌ Database URL is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}

	DB = db
	log.Println("✅ Database connected successfully")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	fmt.Println("✅ Migration successful")
}
