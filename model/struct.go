package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" gorm:"unique"`
	Lastname   string `json:"lastname"`
	Age        uint   `json:"age"`
	Password   string `json:"password"`
	UserRole   string `json:"role"`
	Books      []Book
}

type Book struct {
	gorm.Model      `swaggerignore:"true"`
	Name            string `json:"name" gorm:"unique"`
	Gender          string `json:"gender"`
	Yearpublication string `json:"yearpublication"`
	Author          string `json:"author"`
	PubCompany      string `json:"pubcompany"`
	Photo           string `json:"photo"`
	UserID          uint   `json:"user_id"`
}

type Request struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
