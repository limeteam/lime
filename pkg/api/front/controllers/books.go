package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
	"lime/pkg/api/front/dto"
	"lime/pkg/api/front/service"
	"lime/pkg/api/utils/e"
)

type BooksController struct {
	controllers.BaseController
}
var BooksService = service.BooksService{}

func (C *BooksController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data, err := BooksService.GetBookInfoById(Dto)
		if err != nil {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}