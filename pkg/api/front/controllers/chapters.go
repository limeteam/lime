package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
	"lime/pkg/api/front/dto"
	"lime/pkg/api/front/service"
	"lime/pkg/api/utils/e"
)

type ChaptersController struct {
	controllers.BaseController
}
var ChaptersService = service.ChaptersService{}

func (C *ChaptersController) List(c *gin.Context) {
	var Dto dto.ChapterListDto
	if C.BindAndValidate(c, &Dto) {
		data := ChaptersService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *ChaptersController) Get(c *gin.Context) {
	var Dto dto.ChaptersGetDto
	if C.BindAndValidate(c, &Dto) {
		data,err := ChaptersService.GetChapterInfoById(Dto)
		if err != nil {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}