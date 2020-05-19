package router

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/controllers"
	"lime/pkg/api/middleware"
	"lime/pkg/api/utils/upload"
	"net/http"
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
	e.GET("/novels/categories", CategoryController.List)
	e.GET("/novels/categories/:id", CategoryController.Get)
	e.POST("/novels/categories", CategoryController.Create)
	e.PUT("/novels/categories/:id", CategoryController.Edit)
	e.DELETE("/novels/categories/:id", CategoryController.Delete)

	BooksController := &controllers.BooksController{}
	e.GET("/novels/books", BooksController.List)
	e.GET("/novels/books/:id", BooksController.Get)
	e.POST("/novels/books", BooksController.Create)
	e.PUT("/novels/books/:id", BooksController.Edit)
	e.DELETE("/novels/books/:id", BooksController.Delete)
	e.POST("/novels/books/:id/status", BooksController.UpdateStatus)

	e.POST("/upload/cover", BooksController.UploadBookCover)    // 上传封面图片
	e.POST("/upload/book_file", BooksController.UploadBookFile) // 上传书本

	ChaptersController := &controllers.ChaptersController{}
	e.GET("/novels/chapters", ChaptersController.List)
	e.GET("/novels/chapters/:id", ChaptersController.Get)
	e.POST("/novels/chapters/:id", ChaptersController.Edit)
	e.DELETE("/novels/chapters/:id", ChaptersController.Delete)

	CommentsController := &controllers.CommentsController{}
	e.GET("/novels/comments", CommentsController.List)
	e.GET("/novels/comments/:id", CommentsController.Get)
	e.DELETE("/novels/comments/:id", CommentsController.Delete)

	// 漫画中心

	ComicCategoryController := &controllers.ComicCategoryController{}
	e.GET("/comics/categories", ComicCategoryController.List)
	e.GET("/comics/categories/:id", ComicCategoryController.Get)
	e.POST("/comics/categories", ComicCategoryController.Create)
	e.PUT("/comics/categories/:id", ComicCategoryController.Edit)
	e.DELETE("/comics/categories/:id", ComicCategoryController.Delete)

	//ComicChaptersController := &controllers.ComicChaptersController{}
	//e.GET("/comics/chapters", ComicChaptersController.List)
	//e.GET("/comics/chapters/:id", ComicChaptersController.Get)
	//e.POST("/comics/chapters", ComicChaptersController.Create)
	//e.PUT("/comics/chapters/:id", ComicChaptersController.Edit)
	//e.DELETE("/comics/chapters/:id", ComicChaptersController.Delete)


	e.LoadHTMLGlob("./pkg/ui/dist/*.html")              // 添加入口index.html
	e.LoadHTMLFiles("./pkg/ui/dist/static/*/*")         // 添加资源路径
	e.Static("/static", "./pkg/ui/dist/static")         // 添加资源路径
	e.StaticFile("/admin/", "./pkg/ui/dist/index.html") //前端接口
	e.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	e.StaticFS("/upload/books", http.Dir("data/books/"))
}
