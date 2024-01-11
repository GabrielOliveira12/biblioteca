package controller

import (
	"biblioteca/db"
	"biblioteca/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func InsereUsuario(c *gin.Context) {
	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	usuario.Senha = hashedSenha

	result := db.Database.Create(&usuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, usuario)
}

func ListaUsuario(c *gin.Context) {
	authenticatedUserID, _ := c.Get("UserID")

	var usuario []model.Usuario
	result := db.Database.Where("id = ?", authenticatedUserID).Find(&usuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, usuario)
}
func EditaUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuario model.Livro
	if err := db.Database.First(&usuario, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "usuario não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Database.Save(&usuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, usuario)
}

func DeletaUsuario(c *gin.Context) {
	id := c.Param("id")

	result := db.Database.Delete(&model.Usuario{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuário excluído com sucesso"})
}
