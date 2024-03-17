package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roman91DE/budgetTracker/database"
	"github.com/roman91DE/budgetTracker/handlers"
)

func main() {
	// Initialize DB
	db := database.InitDB("main.db")

	// Load JWT secret from environment variable or configuration file
	jwtSecret := []byte("your_secret_key")

	r := gin.Default()

	// Define routes
	r.GET("/", handlers.GetIndex)
	r.POST("/login", handlers.PostLogin(db, jwtSecret))
	r.POST("/register", handlers.PostRegister(db))
	r.GET("/users", handlers.JWTAuthMiddleware(jwtSecret), handlers.GetUsers(db))

	r.Run() // listens on port 8080 by default
}
