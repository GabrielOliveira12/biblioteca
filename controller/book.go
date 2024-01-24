package controller

import (
	"biblioteca/db"
	"biblioteca/model"

	"github.com/gin-gonic/gin"
)

// InsertBook godoc
// @Summary Insert Book
// @Schemes
// @Param Books body model.Book true "Book structure"
// @Description Insert Books
// @Tags Books
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
// @Router /books [post]
func InsertBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authenticatedUserID, _ := c.Get("UserID")

	userID, ok := authenticatedUserID.(float64)
	if !ok {
		c.JSON(500, gin.H{"error": "Error converting UserID to uint"})
		return
	}
	book.UserID = uint(userID)

	result := db.Database.Create(&book)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error inserting the book", "details": result.Error.Error()})
		return
	}

	c.JSON(200, book)
}

// ListBook godoc
// @Summary List Books
// @Schemes
// @Tags Books
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
// @Router /books [get]
func ListBook(c *gin.Context) {
	var books []model.Book

	authenticatedUserID, _ := c.Get("UserID")
	userID, ok := authenticatedUserID.(float64)
	if !ok {
		c.JSON(500, gin.H{"error": "Error converting UserID to uint"})
		return
	}

	result := db.Database.Where("user_id = ?", uint(userID)).Find(&books)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, books)
}

// UpdateBook godoc
// @Summary Update Book
// @Schemes
// @Param Books body model.Book true "Book structure"
// @Param id path string true "Book ID"
// @Description Update Books
// @Tags Books
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
// @Router /books/{id}  [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book model.Book
	if err := db.Database.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	authenticatedUserID, _ := c.Get("UserID")
	userID, ok := authenticatedUserID.(float64)
	if !ok {
		c.JSON(500, gin.H{"error": "Error converting UserID to uint"})
		return
	}

	if book.UserID != uint(userID) {
		c.JSON(403, gin.H{"error": "Permission denied. You can only edit books linked to your own profile."})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Save(&book)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, book)
}

// DeleteBook godoc
// @Summary Delete Book
// @Schemes
// @Param id path string true "Book ID"
// @Description Delete Books
// @Tags Books
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
// @Router /books/{id}  [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book model.Book
	if err := db.Database.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	authenticatedUserID, _ := c.Get("UserID")
	userID, ok := authenticatedUserID.(float64)
	if !ok {
		c.JSON(500, gin.H{"error": "Error converting UserID to uint"})
		return
	}

	if book.UserID != uint(userID) {
		c.JSON(403, gin.H{"error": "Permission denied. You can only delete books linked to your own profile."})
		return
	}

	result := db.Database.Delete(&model.Book{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Book deleted successfully"})
}
