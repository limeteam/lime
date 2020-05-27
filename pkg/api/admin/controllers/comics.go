package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var ComicsService = service.ComicsService{}

type ComicsController struct {
	BaseController
}

func (C *ComicsController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := ComicsService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (C *ComicsController) Create(c *gin.Context) {
	var Dto dto.ComicsCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := ComicsService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *ComicsController) Edit(c *gin.Context) {
	var Dto dto.ComicsEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicsService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ComicsController) UpdateStatus(c *gin.Context) {
	var Dto dto.ComicsUpdateStatusDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicsService.UpdateStatus(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ComicsController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := ComicsService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *ComicsController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicsService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}

func (C *ComicsController) UploadComicsCover(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	filename := header.Filename
	result, err := ComicsService.UploadCover(file, filename)
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	C.Resp(c, map[string]interface{}{
		"result": result,
	})
}

func (C *ComicsController) UploadComicsFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	filename := header.Filename
	result, err := ComicsService.UploadComicsFile(file, filename)
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	C.Resp(c, map[string]interface{}{
		"result": result,
	})
}
