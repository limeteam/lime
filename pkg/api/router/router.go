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

	taskController := &controllers.TaskController{}
	e.GET("/tasks", taskController.List)
	e.POST("/tasks", taskController.Create)
	e.GET("/tasks/:id/detail", taskController.Get)
	e.PUT("/tasks/:id", taskController.Edit)
	e.PUT("/tasks/:id/stop", taskController.StopTask)
	e.PUT("/tasks/:id/start", taskController.StartTask)
	e.PUT("/tasks/:id/restart", taskController.RestartTask)

	ruleController := &controllers.RuleController{}
	e.GET("/rules", ruleController.GetRuleList)

	e.LoadHTMLGlob("./pkg/ui/dist/*.html")              // 添加入口index.html
	e.LoadHTMLFiles("./pkg/ui/dist/static/*/*")         // 添加资源路径
	e.Static("/static", "./pkg/ui/dist/static")         // 添加资源路径
	e.StaticFile("/admin/", "./pkg/ui/dist/index.html") //前端接口
}
