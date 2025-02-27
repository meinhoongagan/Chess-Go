package controllers

import (
	"Chess-Go/database"
	"Chess-Go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func (a *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// var email = user.Email
	// var username = user.Username

	// isExist := database.DB.Where("email = ?", email).Or("username = ?", username).First(&user)
	// if isExist.Error != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Already exist"})
	// 	return
	// }

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
