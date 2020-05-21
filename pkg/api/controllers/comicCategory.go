package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"lime/pkg/api/utils/e"
)

var ComicCategoryService = service.ComicCategoryService{}

type ComicCategoryController struct {
	BaseController
}

func (C *ComicCategoryController) List(c *gin.Context) {
	var Dto dto.ComicCategoryListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := ComicCategoryService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      Dto.Page,
			"page_size": Dto.Limit,
		})
	}
}
func (C *ComicCategoryController) Create(c *gin.Context) {
	var Dto dto.ComicCategoryCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := ComicCategoryService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *ComicCategoryController) Edit(c *gin.Context) {
	var Dto dto.ComicCategoryEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicCategoryService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ComicCategoryController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := ComicCategoryService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *ComicCategoryController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicCategoryService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}
