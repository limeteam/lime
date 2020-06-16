package router

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
)

func Admin(e *gin.Engine) {
	//后台接口
	admin := e.Group("/admin")
	AdminUsersController := &controllers.AdminUsersController{}
	admin.GET("/users", AdminUsersController.List)
	admin.GET("/users/:id", AdminUsersController.Get)
	admin.POST("/users", AdminUsersController.Create)
	admin.PUT("/users/:id", AdminUsersController.Edit)

	AdminUploadController := &controllers.UploadController{}
	admin.GET("/upload/qiniuToken", AdminUploadController.QiniuToken) //七牛token

	CategoryController := &controllers.CategoryController{}
	admin.GET("/novels/categories", CategoryController.List)
	admin.GET("/novels/categories/:id", CategoryController.Get)
	admin.POST("/novels/categories", CategoryController.Create)
	admin.PUT("/novels/categories/:id", CategoryController.Edit)
	admin.DELETE("/novels/categories/:id", CategoryController.Delete)

	BooksController := &controllers.BooksController{}
	admin.GET("/novels/books", BooksController.List)
	admin.GET("/novels/books/:id", BooksController.Get)
	admin.POST("/novels/books", BooksController.Create)
	admin.PUT("/novels/books/:id", BooksController.Edit)
	admin.DELETE("/novels/books/:id", BooksController.Delete)
	admin.POST("/novels/books/:id/status", BooksController.UpdateStatus)

	admin.POST("/upload/cover", BooksController.UploadBookCover)    // 上传封面图片
	admin.POST("/upload/book_file", BooksController.UploadBookFile) // 上传书本

	ChaptersController := &controllers.ChaptersController{}
	admin.GET("/novels/chapters", ChaptersController.List)
	admin.GET("/novels/chapters/:id", ChaptersController.Get)
	admin.POST("/novels/chapters/:id", ChaptersController.Edit)
	admin.DELETE("/novels/chapters/:id", ChaptersController.Delete)

	CommentsController := &controllers.CommentsController{}
	admin.GET("/novels/comments", CommentsController.List)
	admin.GET("/novels/comments/:id", CommentsController.Get)
	admin.DELETE("/novels/comments/:id", CommentsController.Delete)

	// 漫画中心

	ComicCategoryController := &controllers.ComicCategoryController{}
	admin.GET("/comics/categories", ComicCategoryController.List)
	admin.GET("/comics/categories/:id", ComicCategoryController.Get)
	admin.POST("/comics/categories", ComicCategoryController.Create)
	admin.PUT("/comics/categories/:id", ComicCategoryController.Edit)
	admin.DELETE("/comics/categories/:id", ComicCategoryController.Delete)

	ComicsController := &controllers.ComicsController{}
	admin.GET("/comics/comic", ComicsController.List)
	admin.GET("/comics/comic/:id", ComicsController.Get)
	admin.POST("/comics/comic", ComicsController.Create)
	admin.PUT("/comics/comic/:id", ComicsController.Edit)
	admin.DELETE("/comics/comic/:id", ComicsController.Delete)
	admin.POST("/comics/comic/:id/status", ComicsController.UpdateStatus)


	ComicCommentsController := &controllers.ComicCommentsController{}
	admin.GET("/comics/comments", ComicCommentsController.List)
	admin.GET("/comics/comments/:id", ComicCommentsController.Get)
	admin.DELETE("/comics/comments/:id", ComicCommentsController.Delete)

	//ComicChaptersController := &controllers.ComicChaptersController{}
	//e.GET("/comics/chapters", ComicChaptersController.List)
	//e.GET("/comics/chapters/:id", ComicChaptersController.Get)
	//e.POST("/comics/chapters", ComicChaptersController.Create)
	//e.PUT("/comics/chapters/:id", ComicChaptersController.Edit)
	//e.DELETE("/comics/chapters/:id", ComicChaptersController.Delete)
}
