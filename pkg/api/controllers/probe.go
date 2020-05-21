package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/utils/e"
)

type ProbeController struct {
	BaseController
}

// @Tags Health
// @Summary 健康检查
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{""}}"
// @Router /healthcheck [get]
func (C *ProbeController) Healthy(c *gin.Context) {
	//resp(c, gin.H{
	//	"data":   true,
	//})
	C.Fail(c, e.Err404)
}
