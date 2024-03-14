package db

import (
	"github.com/roman91DE/budgetTracker/models"
	"os"
	"path/filepath"
	"testing"
)

// TestInitDB tests the initialization and migration of the database.
func TestInitDB(t *testing.T) {
	// Create a temporary directory for the test database.
	tempDir, err := os.MkdirTemp("", "testdb")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test.

	// Construct the path to the temporary database file.
	tempDBPath := filepath.Join(tempDir, "test.db")

	// Initialize the database with the temporary file.
	db := InitDB(tempDBPath)
	if db == nil {
		t.Fatal("Expected db to be initialized, got nil")
	}

	// Use GORM's migration feature to check if the `User` table exists by attempting to migrate it.
	// This is a rudimentary check to ensure migrations are applied.
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}
}
