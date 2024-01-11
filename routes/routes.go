package routes

import (
	"biblioteca/auth"
	"biblioteca/controller"
	"biblioteca/middleware"

	"github.com/gin-gonic/gin"
)

func ApiBiblioteca(r *gin.Engine) {

	openuser := r.Group("usuarios")
	privatyuser := r.Group("usuarios")
	admprivatyuser := r.Group("livros")

	openuser.Use(middleware.OpenUserMiddleware())
	privatyuser.Use(middleware.PrivatyUserMiddleware())
	admprivatyuser.Use(auth.TokenAuthMiddleware())

	openuser.POST("", controller.InsereUsuario)
	privatyuser.GET("", controller.ListaUsuario)
	privatyuser.PUT("/:id", controller.EditaUsuario)
	privatyuser.DELETE("/:id", controller.DeletaUsuario)

	admprivatyuser.POST("", controller.InsereLivro)
	admprivatyuser.GET("", controller.ListaLivro)
	admprivatyuser.PUT("/:id", controller.EditarLivro)
	admprivatyuser.DELETE("/:id", controller.DeletaLivro)

	r.POST("/requests", controller.Login)

}
