package main

import (
	"biblioteca/db"
	"biblioteca/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Connect()
	routes.ApiBiblioteca(r)
	r.Run(":8080")
}
