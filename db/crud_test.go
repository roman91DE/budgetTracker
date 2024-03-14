package db

import (
	"github.com/roman91DE/budgetTracker/models"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateAndDeleteUser(t *testing.T) {

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

	// Test CreateUserInDB
	testEmail := "test@example.com"
	testPassword := "password123"
	user, _ := models.MakeUser(testEmail, testPassword)
	err = CreateUserInDB(user, db)
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}

	// Verify user was created
	var count int64
	db.Model(&models.User{}).Where("email = ?", testEmail).Count(&count)
	if count != 1 {
		t.Errorf("Expected 1 user, found %d", count)
	}

	// Verify email must be unique
	err = CreateUserInDB(user, db)
	if err == nil {
		t.Errorf("Created User with not-unique E-Mail: %v", err)
	}

	// Test DeleteUserInDB
	deleted, err := DeleteUserInDB(testEmail, db)
	if err != nil {
		t.Errorf("Failed to delete user: %v", err)
	}
	if !deleted {
		t.Errorf("Expected user to be deleted, but was not")
	}

	// Verify user was deleted
	db.Model(&models.User{}).Where("email = ?", testEmail).Count(&count)
	if count != 0 {
		t.Errorf("Expected 0 users, found %d", count)
	}
	// we should still be unable to recreate a user because of gorms soft delete feature
	err = CreateUserInDB(user, db)
	if err == nil {
		t.Errorf("Reused a unique identifier for E-Mail: %v", err)
	}
}
