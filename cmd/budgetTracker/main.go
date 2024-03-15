package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roman91DE/budgetTracker/database"
	"github.com/roman91DE/budgetTracker/models"
	"net/http"
)

func main() {

	r := gin.Default()
	db := database.InitDB("main.db")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, budgetTracker!",
		})
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
