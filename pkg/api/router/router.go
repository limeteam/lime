package router

import (
	"github.com/gin-gonic/gin"
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

	e.LoadHTMLGlob("./pkg/ui/dist/*.html")              // 添加入口index.html
	e.LoadHTMLFiles("./pkg/ui/dist/static/*/*")         // 添加资源路径
	e.Static("/static", "./pkg/ui/dist/static")         // 添加资源路径
	e.StaticFile("/admin/", "./pkg/ui/dist/index.html") //前端接口
	e.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	e.StaticFS("/upload/books", http.Dir("data/books/"))
}
