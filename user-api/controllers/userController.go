// controllers/userController.go

package controllers

import (
	"fmt"
	"net/http"
	"user-reservation-api/dtos"
	"user-reservation-api/initializers"
	"user-reservation-api/models"
	"user-reservation-api/services"

	"github.com/gin-gonic/gin"
)

// SignUp controller function
func SignUp(c *gin.Context) {
	services.SignUp(c)
}

// Login controller function
func Login(c *gin.Context) {
	var dto dtos.LoginUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := services.Login(dto, c) // Pasar dto por valor en lugar de referencia
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Validate controller function
func Validate(c *gin.Context) {
	user, err := services.Validate(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var foundUser models.User
	if err := initializers.DB.First(&foundUser, user.(models.User).ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": foundUser})
}

func Logout(c *gin.Context) {
	// Eliminar la cookie de autorización
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// CheckUserExistence verifica si un usuario existe
func CheckUserExistence(c *gin.Context) {
	userID := c.Param("userID")

	var user models.User
	if err := initializers.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		// Si el usuario no existe, devolver 404
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with ID %s does not exist", userID)})
			return
		}
		// Error en la base de datos
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Si el usuario existe, devolver 200 OK
	c.JSON(http.StatusOK, gin.H{"message": "User exists"})
}
