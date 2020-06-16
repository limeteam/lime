package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lime/pkg/api/admin/service"
)

var UploadService = service.UploadService{}

type UploadController struct {
	BaseController
}

func (C *UploadController) QiniuToken(c *gin.Context) {
	token := UploadService.GetToken()
	C.Resp(c, map[string]interface{}{
		"token": token,
		"domain": viper.GetString("qiniu.domain"),
	})
}
