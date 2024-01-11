package model

import (
	"gorm.io/gorm"
)

type Livro struct {
	gorm.Model
	Nome          string `json:"nome" gorm:"unique"`
	Genero        string `json:"genero"`
	AnoPublicacao string `json:"anopublicacao"`
	Autor         string `json:"autor"`
	Editora       string `json:"editora"`
	Foto          string `json:"foto"`
}

type Usuario struct {
	gorm.Model
	Nome      string `json:"nome" gorm:"unique"`
	Sobrenome string `json:"sobrenome"`
	Idade     uint   `json:"idade"`
	Senha     []byte `json:"senha"`
	UserRole  string `json:"role"`
}

type Request struct {
	Nome  string `json:"nome" binding:"required"`
	Senha []byte `json:"senha" binding:"required"`
}
