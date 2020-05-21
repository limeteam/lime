package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"lime/pkg/api/utils/e"
)

var CategoryService = service.CategoryService{}

type CategoryController struct {
	BaseController
}

func (C *CategoryController) List(c *gin.Context) {
	var CategoryListDto dto.CategoryListDto
	if C.BindAndValidate(c, &CategoryListDto) {
		data, total := CategoryService.List(CategoryListDto)
		C.Resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      CategoryListDto.Page,
			"page_size": CategoryListDto.Limit,
		})
	}
}
func (C *CategoryController) Create(c *gin.Context) {
	var Dto dto.CategoryCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := CategoryService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *CategoryController) Edit(c *gin.Context) {
	var Dto dto.CategoryEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := CategoryService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *CategoryController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := CategoryService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *CategoryController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := CategoryService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}
