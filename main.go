package main

import (
	"biblioteca/db"
	"biblioteca/docs"
	"biblioteca/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Library
//	@version		1.0.0
//	@description	Library
//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
//
// @host	localhost:8080
// @BasePath	/
// @schemes	http
func main() {
	r := gin.Default()
	db.Connect()
	routes.ApiLibrary(r)

	docs.SwaggerInfo.BasePath = ""

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8080")

}
