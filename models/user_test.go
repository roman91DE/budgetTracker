// models/user_test.go

package models

import "testing"

// TestUserInitialization tests the initialization of a User struct.
func TestUserInitialization(t *testing.T) {
    user := User{
        ID:       1,
        Email:    "test@example.com",
        Password: "password123",
    }

    // Check if the User struct is initialized with the correct values.
    if user.ID != 1 {
        t.Errorf("Expected ID to be '1', got %v", user.ID)
    }
    if user.Email != "test@example.com" {
        t.Errorf("Expected Email to be 'test@example.com', got %v", user.Email)
    }
    if user.Password != "password123" {
        t.Errorf("Expected Password to be 'password123', got %v", user.Password)
    }
}

func
