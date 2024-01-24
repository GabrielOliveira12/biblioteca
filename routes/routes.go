package routes

import (
	"biblioteca/controller"
	"biblioteca/middleware"

	"github.com/gin-gonic/gin"
)

func ApiLibrary(r *gin.Engine) {

	openuser := r.Group("users")
	privatyuser := r.Group("users")
	admprivatyuser := r.Group("books")

	openuser.Use(middleware.OpenUserMiddleware())
	privatyuser.Use(middleware.PrivatyUserMiddleware())
	admprivatyuser.Use(middleware.AdmPrivatyUserMiddleware())

	openuser.POST("", controller.InsertUser)
	privatyuser.GET("", controller.ListUser)
	privatyuser.PUT("/:id", controller.UpdateUser)
	privatyuser.DELETE("/:id", controller.DeleteUser)

	admprivatyuser.POST("", controller.InsertBook)
	admprivatyuser.GET("", controller.ListBook)
	admprivatyuser.PUT("/:id", controller.UpdateBook)
	admprivatyuser.DELETE("/:id", controller.DeleteBook)

	r.POST("/requests", controller.Login)

}
