package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lime/pkg/api/utils"
	"strings"
)

type SettingController struct {
	BaseController
}

func (C *SettingController) GetTokenAndEncodingAESKey(c *gin.Context) {
	C.Resp(c, map[string]interface{}{
		"receiveUrl": viper.GetString("wechat.receiveUrl") + "/v1/wechat/callback",
		"token": strings.ToLower(utils.CreateRandomString(32)),
		"EncodingAESKey": utils.CreateRandomString(43),
	})
}