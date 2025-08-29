package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	var err error
	var db *gorm.DB

	// Retry connection with backoff for Docker startup
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * 2 * time.Second)
		}
	}

	if err != nil {
		log.Fatal("Failed to connect to database after retries:", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto-migrate the schema
	err = db.AutoMigrate(&User{}, &Portfolio{}, &Asset{}, &AssetHistory{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Database connected and migrated successfully")
}
