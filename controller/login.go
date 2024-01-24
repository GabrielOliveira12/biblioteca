package controller

import (
	"biblioteca/auth"
	"biblioteca/db"
	"biblioteca/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login
// @Schemes
// @Param request body model.Request true "Login credentials"
// @Description Handles user login
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {string} Login
// @Router /requests [post]
func Login(c *gin.Context) {
	var request model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := db.Database.Where("name = ?", request.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	passwordBytes := []byte(request.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), passwordBytes); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
