package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
)

type ChaptersController struct {
	controllers.BaseController
}
var ChaptersService = service.ChaptersService{}

func (C *ChaptersController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data := ChaptersService.FrontList(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *ChaptersController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := ChaptersService.GetChapterInfoById(Dto)
		//if data.Id < 1 {
		//	C.Fail(c, e.ErrIdData)
		//	return
		//}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}