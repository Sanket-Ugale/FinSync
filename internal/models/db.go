package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{}, &Portfolio{}, &Asset{}, &AssetHistory{})

	DB = db
}
