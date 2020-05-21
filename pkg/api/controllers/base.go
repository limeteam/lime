package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/utils/e"
	"net/http"
)

type BaseController struct {
}

func (bc *BaseController) BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := dto.Bind(c, obj); err != nil {
		failValidate(c, err.Error())
		return false
	}
	return true
}

func (bc *BaseController) Resp(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func (bc *BaseController) Ok(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
}

func (bc *BaseController) Fail(c *gin.Context, errs *e.ControllerError) {
	c.JSON(http.StatusOK, gin.H{
		"code":     errs.Code,
		"msg":      errs.Message,
		"moreinfo": errs.Moreinfo,
	})
}

func failValidate(c *gin.Context, msg string) {
	errs := e.ErrValidation
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":   errs.Code,
		"msg":    errs.Message,
		"detail": msg,
	})
}
