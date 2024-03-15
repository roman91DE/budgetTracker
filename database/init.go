package database

import (
	"github.com/roman91DE/budgetTracker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

func InitDB(dbFilePath string) *gorm.DB {
	filepath.Join("db", dbFilePath)
	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	return db
}
