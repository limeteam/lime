package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
	"lime/pkg/api/front/service"
)

type WechatController struct {
	controllers.BaseController
}
var WechatService = service.WechatService{}

func (C *WechatController) Callback(c *gin.Context) {
	WechatService.Callback(c)
}