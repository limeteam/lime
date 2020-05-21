package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"lime/pkg/api/utils/e"
)

var ChaptersService = service.ChaptersService{}

type ChaptersController struct {
	BaseController
}

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
func (C *ChaptersController) Create(c *gin.Context) {
	var Dto dto.ChaptersCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := ChaptersService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *ChaptersController) Edit(c *gin.Context) {
	var Dto dto.ChaptersEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := ChaptersService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ChaptersController) UpdateStatus(c *gin.Context) {
	var Dto dto.ChaptersUpdateStatusDto
	if C.BindAndValidate(c, &Dto) {
		affected := ChaptersService.UpdateStatus(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
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

func (C *ChaptersController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := ChaptersService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}
