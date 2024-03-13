package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

func makeUser(id uint, email, passwd string) (*User, error) {

	if len([]byte(passwd)) > 72 {
		return nil, fmt.Errorf("password is too long! Cannot create user")
	}
	user := &User{
		ID:       id,
		Email:    email,
		Password: passwd,
	}

	user.HashPassword()

	return user, nil
}

// HashPassword generates a hashed password from a plaintext string
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword compares a plaintext password with the user's hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
