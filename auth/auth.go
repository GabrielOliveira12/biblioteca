package auth

import (
	"biblioteca/model"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_key"
	}
	return []byte(secret)
}

func getJWTExpiryHours() time.Duration {
	hours := os.Getenv("JWT_EXPIRY_HOURS")
	if hours == "" {
		return time.Hour * 24
	}
	h, err := strconv.Atoi(hours)
	if err != nil {
		return time.Hour * 24
	}
	return time.Hour * time.Duration(h)
}

func getTokenPrefix() string {
	prefix := os.Getenv("JWT_TOKEN_PREFIX")
	if prefix == "" {
		return "Bearer"
	}
	return prefix
}

func GenerateJWTToken(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["role"] = user.UserRole
	claims["exp"] = time.Now().Add(getJWTExpiryHours()).Unix()

	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := ExtractTokenFromRequest(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("UserID", claims["id"])
		c.Set("UserRole", claims["role"])

	}
}

func ExtractTokenFromRequest(req *http.Request) (string, error) {
	authorizationHeader := req.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("Token não encontrado")
	}

	tokenParts := strings.Split(authorizationHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != getTokenPrefix() {
		return "", errors.New("Formato inválido de token")
	}

	return tokenParts[1], nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Método de assinatura inválido")
		}

		return getJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Erro ao obter reivindicações do token")
	}

	return claims, nil
}
