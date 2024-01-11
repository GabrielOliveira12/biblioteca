package middleware

import (
	"biblioteca/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OpenUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func PrivatyUserMiddleware() gin.HandlerFunc {
	return auth.TokenAuthMiddleware()
}

func AdmPrivatyUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenAuthMiddleware()(c)
		UserRole, exists := c.Get("UserRole")
		if !exists || UserRole.(string) != "adm" {

			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado. Função de administrador necessária."})
			c.Abort()
			return
		}

		c.Next()
	}
}
