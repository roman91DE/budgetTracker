package models

import (
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID       uint   `json:"id" gorm:"primary_key"`
    Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
    Password string `json:"password"`
}

// HashPassword generates a hashed password from a plaintext string
func (u *User) HashPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
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
