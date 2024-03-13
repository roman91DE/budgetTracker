// models/user_test.go

package models

import (
	"crypto/rand"
	"encoding/base64"
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
}

func TestPasswordHashing(t *testing.T) {

	plaintextPassword := "password123"

	user, _ := makeUser(1, "test@example.com", plaintextPassword)

	if user.Password == plaintextPassword {
		t.Errorf("Expected hashed Password but still the same as Plaintext Argument")
	}
}

func TestPasswordChecks(t *testing.T) {

	plaintextPassword := "password123"

	user, _ := makeUser(1, "test@example.com", plaintextPassword)

	falsePassword := "321password"

	if user.CheckPassword(falsePassword) {
		t.Errorf("False Password %v matched the correct Passord %v", falsePassword, plaintextPassword)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plaintextPassword))
	if err != nil {
		t.Errorf("Correct plaintext Password %v did not match: %v", plaintextPassword, err)
	}

}

func TestFailOnLongPasswords(t *testing.T) {

	// Generate 73 random bytes
	randomBytes := make([]byte, 73)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Handle the error here
		t.Errorf("Unrelated Error during random Password Generation: %v", err)
	}

	// Encode the bytes into a base64 string
	randomLongString := base64.StdEncoding.EncodeToString(randomBytes)

	_, err = makeUser(1, "test@example.com", randomLongString)

	if err == nil {
		t.Errorf("Initialized user with invalid, too long password: %v", err)
	}

}
