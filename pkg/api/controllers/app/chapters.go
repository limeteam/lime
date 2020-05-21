package app

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/controllers"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"lime/pkg/api/utils/e"
)

type ChaptersController struct {
	controllers.BaseController
}
var ChaptersService = service.ChaptersService{}

func (C *ChaptersController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := ChaptersService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}

func (C *ChaptersController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := ChaptersService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}