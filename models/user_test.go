// models/user_test.go

package models

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

// TestUserInitialization tests the initialization of a User struct.
func TestUserInitialization(t *testing.T) {
	user, err := makeUser(1, "test@example.com", "password123")

	if err != nil {
		t.Errorf("makeUser failed on valid input: %v", err)
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

func TestPasswordHashing(t *testing.T) {

	plaintextPassword := "password123"
	user := User{
		ID:       1,
		Email:    "test@example.com",
		Password: plaintextPassword,
	}

	user.HashPassword()

	if user.Password == plaintextPassword {
		t.Errorf("Expected hashed Password but still the same as Plaintext Argument")
	}
}

func TestPasswordChecks(t *testing.T) {

	plaintextPassword := "password123"
	user := User{
		ID:       1,
		Email:    "test@example.com",
		Password: plaintextPassword,
	}

	user.HashPassword()

	falsePassword := "321password"

	if user.CheckPassword(falsePassword) {
		t.Errorf("False Password %v matched the correct Passord %v", falsePassword, plaintextPassword)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plaintextPassword))
	if err != nil {
		t.Errorf("Correct plaintext Password %v did not match: %v", plaintextPassword, err)
	}

}
