package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/roman91DE/budgetTracker/database"
	"github.com/roman91DE/budgetTracker/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

func main() {

	r := gin.Default()
	db := database.InitDB("main.db")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, budgetTracker!",
		})
	})

	r.POST("/login", func(c *gin.Context) {

		var req models.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(req.Password)

		// Check if the email exists in the database
		if !database.EmailExists(req.Email, db) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is note registered"})
			return
		}

		// check if the password matches the hash
		ok, hash := database.GetHashedPassword(req.Password, db)
		fmt.Println(hash)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could fetch hashed Password from database!"})
			return
		}
		// Compare the hash with the password provided by the user
        err := bcrypt.CompareHashAndPassword(hash, []byte(req.Password))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password!"})
            return
        }

		// Create JWT token
		expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
		claims := &models.Claims{
			Email: req.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Return the token in the response
		c.JSON(http.StatusOK, models.LoginResponse{Token: tokenString})
	})

	r.POST("/register", func(c *gin.Context) {
		var req models.RegisterRequest

		// Parse JSON body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Check if the email is already registered
		var count int64
		db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count)
		if count > 0 {
			// email is already in use
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered"})
			return
		}

		user, err := models.MakeUser(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = database.CreateUserInDB(user, db); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Registered new User": user.Email})
	})

	r.GET("/users", func(c *gin.Context) {

		emails, err := database.GetUserbase(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// If successful, respond with the list of emails
		c.JSON(http.StatusOK, gin.H{"emails": emails})
	})

	r.Run() // By default, it listens on port 8080
}
