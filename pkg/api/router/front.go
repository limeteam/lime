package router

import (
	"github.com/gin-gonic/gin"
	controllersFront "lime/pkg/api/front/controllers"
)

func Front(e *gin.Engine) {
	//前台接口
	v1 := e.Group("/v1")
	UsersController := &controllersFront.UsersController{}
	v1.POST("/users",UsersController.Login)

	V1BooksController := &controllersFront.BooksController{}
	//v1.GET("/books", V1BooksController.List)
	v1.GET("/books/:id", V1BooksController.Get)

	V1ChaptersController := &controllersFront.ChaptersController{}
	v1.GET("/chapters", V1ChaptersController.List)
	v1.GET("/chapters/:id", V1ChaptersController.Get)
}
