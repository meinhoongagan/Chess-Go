package controllers

import (
	"Chess-Go/database"
	"Chess-Go/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Secret key for signing JWT (store securely)
var jwtSecret = []byte("your_secret_key")

type AuthController struct{}

// Function to generate JWT token
func generateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (a *AuthController) Login(c *gin.Context) {
	var userInput models.User
	fmt.Println("Login request received")

	// Bind JSON request body
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("User Data:", userInput)

	// Check if user exists by email OR username
	var existingUser models.User
	emailExists := database.DB.Where("email = ?", userInput.Email).First(&existingUser)
	usernameExists := database.DB.Where("username = ?", userInput.Username).First(&existingUser)

	if emailExists.Error == nil && usernameExists.Error == nil {
		// Email and username both match -> Allow login
		token, err := generateToken(existingUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
		return
	}

	if emailExists.Error == nil || usernameExists.Error == nil {
		// Only email or username exists but not both -> Reject
		c.JSON(http.StatusConflict, gin.H{"error": "Email or username already exists"})
		return
	}

	// Neither email nor username exist -> Create new user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashedPassword)

	// Save new user in DB
	result := database.DB.Create(&userInput)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Generate JWT token for the new user
	token, err := generateToken(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": token})
}
