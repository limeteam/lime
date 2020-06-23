package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var ComicChaptersService = service.ComicChaptersService{}

type ComicChaptersController struct {
	BaseController
}

func (C *ComicChaptersController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := ComicChaptersService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (C *ComicChaptersController) Create(c *gin.Context) {
	var Dto dto.ComicChaptersCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := ComicChaptersService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *ComicChaptersController) Edit(c *gin.Context) {
	var Dto dto.ComicChaptersEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicChaptersService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ComicChaptersController) UpdateStatus(c *gin.Context) {
	var Dto dto.ComicChaptersUpdateStatusDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicChaptersService.UpdateStatus(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ComicChaptersController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := ComicChaptersService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *ComicChaptersController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicChaptersService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}
