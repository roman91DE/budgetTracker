package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/roman91DE/budgetTracker/database"
	"github.com/roman91DE/budgetTracker/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, budgetTracker!",
	})
}

// GenerateJWTToken generates a JWT token for the given email
func GenerateJWTToken(email string, jwtSecret []byte) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	// Create JWT claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   email,
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


// PostLogin handles user login
func PostLogin(db *gorm.DB, jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Check if the email exists in the database
		if !database.EmailExists(req.Email, db) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is not registered"})
			return
		}

		// Check if the password matches the hash
		ok, hash := database.GetHashedPassword(req.Email, db)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch hashed password from database"})
			return
		}

		if err := bcrypt.CompareHashAndPassword(hash, []byte(req.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
			return
		}

		// Generate JWT token
		tokenString, err := GenerateJWTToken(req.Email, jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Return the token in the response
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

// PostRegister handles user registration
func PostRegister(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.RegisterRequest

		// Parse JSON body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Check if the email is already registered
		if database.EmailExists(req.Email, db) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered"})
			return
		}

		// Create user and insert into database
		user, err := models.MakeUser(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		if err := database.CreateUserInDB(user, db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}
}

// GetUsers returns a list of users
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		emails, err := database.GetUserbase(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}

		// If successful, respond with the list of emails
		c.JSON(http.StatusOK, gin.H{"emails": emails})
	}
}
