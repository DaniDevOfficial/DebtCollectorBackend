package config

import (
	"dept-collector/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ failed to connect to database:", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Semester{},
		&models.Amount{},
		&models.Class{},
		&models.Lesson{},
		&models.SkipEntry{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ AutoMigration complete!")
}

func connectMockDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to mock DB:", err)
	}
	log.Println("Using in-memory SQLite DB (mock mode)")
	return db
}
