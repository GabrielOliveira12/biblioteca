package controller

import (
	"biblioteca/db"
	"biblioteca/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// InsertUser godoc
// @Summary Insert User
// @Schemes
// @Param Users body model.User true "User structure"
// @Description Insert Users
// @Tags Users
// @Accept json
// @Produce json
// Authorization @securityDefinitions.apiKey
//
//	@in				header
//	@name			Authorization
//	@Security		JWT
//	@Failure		400 "Bad Request"
//
// @Success 200 {string} InsertUser
// @Router /users [post]
func InsertUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating password hash"})
		return
	}
	hashedPasswordString := string(hashedPassword)

	user.Password = hashedPasswordString

	result := db.Database.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, user)
}

// ListUser godoc
// @Summary List User
// @Schemes
// @Tags Users
// @Accept json
// @Produce json
// Authorization @securityDefinitions.apiKey
//
//	@in				header
//	@name			Authorization
//	@Security		JWT
//	@Success		200 "Success"
//	@Failure		400 "Bad Request"
//
// @Router /users [get]
func ListUser(c *gin.Context) {
	authenticatedUserID, _ := c.Get("UserID")

	var users []model.User
	result := db.Database.Preload("Books").Find(&users, authenticatedUserID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, users)
}

// UpdateUser godoc
// @Summary Update User
// @Schemes
// @Param Books body model.User true "Book structure"
// @Param id path string true "User ID"
// @Description Update Users
// @Tags Users
// @Accept json
// @Produce json
// Authorization @securityDefinitions.apiKey
//
//	@in				header
//	@name			Authorization
//	@Security		JWT
//	@Success		200 "Success"
//	@Failure		400 "Bad Request"
//
// @Router /users/{id}  [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user model.User
	if err := db.Database.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	authenticatedUserID, _ := c.Get("UserID")
	if user.ID != authenticatedUserID {
		c.JSON(403, gin.H{"error": "Permission denied. You can only edit your own profile."})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Save(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, user)
}

// DeleteUser godoc
// @Summary Delete User
// @Schemes
// @Param id path string true "User ID"
// @Description Delete Users
// @Tags Users
// @Accept json
// @Produce json
// Authorization @securityDefinitions.apiKey
//
//	@in				header
//	@name			Authorization
//	@Security		JWT
//	@Failure		400 "Bad Request"
//
// @Success 200 {string} DeleteUser
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user model.User
	if err := db.Database.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	authenticatedUserID, _ := c.Get("UserID")
	if user.ID != authenticatedUserID {
		c.JSON(403, gin.H{"error": "Permission denied. You can only delete your own profile."})
		return
	}

	result := db.Database.Delete(&model.User{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
