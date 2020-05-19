package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
)

var ComicsService = service.ComicsService{}

type ComicsController struct {
	BaseController
}

func (Comics *ComicsController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if Comics.BindAndValidate(c, &Dto) {
		data, total := ComicsService.List(Dto)
		resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (Comics *ComicsController) Create(c *gin.Context) {
	var Dto dto.ComicsCreateDto
	if Comics.BindAndValidate(c, &Dto) {
		created, _ := ComicsService.Create(Dto)
		if created.Id <= 0 {
			fail(c, ErrAddFail)
			return
		}
		resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (Comics *ComicsController) Edit(c *gin.Context) {
	var Dto dto.ComicsEditDto
	if Comics.BindAndValidate(c, &Dto) {
		affected := ComicsService.Update(Dto)
		if affected > 0 {
			ok(c, "更新成功")
			return
		}
		fail(c, ErrEditFail)
		return
	}
}

func (Comics *ComicsController) UpdateStatus(c *gin.Context) {
	var Dto dto.ComicsUpdateStatusDto
	if Comics.BindAndValidate(c, &Dto) {
		affected := ComicsService.UpdateStatus(Dto)
		if affected > 0 {
			ok(c, "更新成功")
			return
		}
		fail(c, ErrEditFail)
		return
	}
}

func (Comics *ComicsController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if Comics.BindAndValidate(c, &Dto) {
		data := ComicsService.InfoOfId(Dto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (Comics *ComicsController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if Comics.BindAndValidate(c, &Dto) {
		affected := ComicsService.Delete(Dto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "删除成功")
	}
}

func (Comics *ComicsController) UploadComicsCover(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		fail(c, ErrUploadCover)
		return
	}
	filename := header.Filename
	result, err := ComicsService.UploadCover(file, filename)
	if err != nil {
		fail(c, ErrUploadCover)
		return
	}
	resp(c, map[string]interface{}{
		"result": result,
	})
}

func (Comics *ComicsController) UploadComicsFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		fail(c, ErrUploadCover)
		return
	}
	filename := header.Filename
	result, err := ComicsService.UploadComicsFile(file, filename)
	if err != nil {
		fail(c, ErrUploadCover)
		return
	}
	resp(c, map[string]interface{}{
		"result": result,
	})
}
