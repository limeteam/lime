package router

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/controllers"
	"lime/pkg/api/middleware"
)

func Init(e *gin.Engine, cors bool) {
	e.Use(
		gin.Recovery(),
	)
	if cors {
		e.Use(middleware.Cors())
	}
	e.GET("/healthcheck", controllers.Healthy)

	CategoryController := &controllers.CategoryController{}
	e.GET("/categories", CategoryController.List)
	e.GET("/categories/:id", CategoryController.Get)
	e.POST("/categories", CategoryController.Create)
	e.PUT("/categories/:id", CategoryController.Edit)
	e.DELETE("/categories/:id", CategoryController.Delete)

	BooksController := &controllers.BooksController{}
	e.GET("/books", BooksController.List)
	e.GET("/books/:id", BooksController.Get)
	e.POST("/books", BooksController.Create)
	e.PUT("/books/:id", BooksController.Edit)
	e.DELETE("/books/:id", BooksController.Delete)

	//CommentsController := &controllers.CommentsController{}
	//e.GET("/comments", CommentsController.List)
	//e.GET("/comments/:id", CommentsController.Get)
	//e.DELETE("/comments/:id", CommentsController.Delete)

	e.LoadHTMLGlob("./pkg/ui/dist/*.html")              // 添加入口index.html
	e.LoadHTMLFiles("./pkg/ui/dist/static/*/*")         // 添加资源路径
	e.Static("/static", "./pkg/ui/dist/static")         // 添加资源路径
	e.StaticFile("/admin/", "./pkg/ui/dist/index.html") //前端接口
}
