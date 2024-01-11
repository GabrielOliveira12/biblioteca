package controller

import (
	"biblioteca/db"
	"biblioteca/model"

	"github.com/gin-gonic/gin"
)

func InsereLivro(c *gin.Context) {
	var livro model.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := db.Database.Create(&livro)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	}

	c.JSON(200, livro)
}

func ListaLivro(c *gin.Context) {
	var livro []model.Livro
	result := db.Database.Find(&livro)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, livro)
}

func EditarLivro(c *gin.Context) {
	id := c.Param("id")

	var livro model.Livro
	if err := db.Database.First(&livro, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "usuario não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Save(&livro)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, livro)
}

func DeletaLivro(c *gin.Context) {
	id := c.Param("id")

	result := db.Database.Delete(&model.Livro{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuário excluído com sucesso"})
}
