package database

import (
	"fmt"
	"github.com/roman91DE/budgetTracker/models"
	"gorm.io/gorm"
)

func CreateUserInDB(user *models.User, db *gorm.DB) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error // Return the error if user creation fails
	}

	return nil // Return nil if user creation is successful
}

func GetUserbase(db *gorm.DB) ([]string, error) {
	var emails []string
	if err := db.Model(&models.User{}).Pluck("Email", &emails).Error; err != nil {
		return nil, fmt.Errorf("couldnt pluck <Email> from database: %v", err)
	}
	return emails, nil
}

func DeleteUserInDB(userEmail string, db *gorm.DB) (bool, error) {
	// Assuming you have a User struct where the Email field is tagged appropriately for GORM
	result := db.Where("email = ?", userEmail).Delete(&models.User{})

	// Check for errors in the delete operation
	if result.Error != nil {
		return false, result.Error
	}

	// Check if any rows were deleted
	if result.RowsAffected == 0 {
		return false, fmt.Errorf("no user found with email %s", userEmail)
	}

	return true, nil
}
