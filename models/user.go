package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

// MakeUser validates arguments, hashes the password and returns a User struct
func MakeUser(email, passwd string) (*User, error) {

	if len([]byte(passwd)) > 72 {
		return nil, fmt.Errorf("password is too long! Cannot create user")
	}

	if len(passwd) < 8 {
		return nil, fmt.Errorf("password is too short! Cannot create user")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(email) {
		return nil, fmt.Errorf("invalid email format! Cannot create user")
	}

	user := &User{
		Email:    email,
		Password: passwd,
	}

	if err := user.hashPassword(); err != nil {
		return nil, fmt.Errorf("error occured during hashing! Cannot create user: %v", err)
	}

	return user, nil
}

// HashPassword generates a hashed password from a plaintext string
func (u *User) hashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword compares a plaintext password with the user's hashed password
func (u *User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}


type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
