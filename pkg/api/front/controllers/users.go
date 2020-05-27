package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
	"lime/pkg/api/admin/dto"
)

type UsersController struct {
	controllers.BaseController
}

func (C *UsersController) Login(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		//data, total := BooksService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": "data",
			"total":  "total",
		})
	}
}
