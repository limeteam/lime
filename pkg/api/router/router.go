package router

import (
	"lime/pkg/api/controllers"
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	e.Use(
		gin.Recovery(),
	)
	//e.Use(cors.Default()) // CORS
	e.GET("/healthcheck", controllers.Healthy)

	CategoryController := &controllers.CategoryController{}
	e.GET("/category/list", CategoryController.List)

	e.LoadHTMLGlob("./pkg/ui/dist/*.html")              // 添加入口index.html
	e.LoadHTMLFiles("./pkg/ui/dist/static/*/*")         // 添加资源路径
	e.Static("/static", "./pkg/ui/dist/static")         // 添加资源路径
	e.StaticFile("/admin/", "./pkg/ui/dist/index.html") //前端接口
}
