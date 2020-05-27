package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var BooksService = service.BooksService{}

type BooksController struct {
	BaseController
}

func (C *BooksController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := BooksService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (C *BooksController) Create(c *gin.Context) {
	var Dto dto.BooksCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := BooksService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *BooksController) Edit(c *gin.Context) {
	var Dto dto.BooksEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := BooksService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *BooksController) UpdateStatus(c *gin.Context) {
	var Dto dto.BooksUpdateStatusDto
	if C.BindAndValidate(c, &Dto) {
		affected := BooksService.UpdateStatus(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *BooksController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := BooksService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *BooksController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := BooksService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}

func (C *BooksController) UploadBookCover(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	filename := header.Filename
	result, err := BooksService.UploadCover(file, filename)
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	C.Resp(c, map[string]interface{}{
		"result": result,
	})
}

func (C *BooksController) UploadBookFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	filename := header.Filename
	result, err := BooksService.UploadBookFile(file, filename)
	if err != nil {
		C.Fail(c, e.ErrUploadCover)
		return
	}
	C.Resp(c, map[string]interface{}{
		"result": result,
	})
}
