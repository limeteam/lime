package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
)

var CategoryService = service.CategoryService{}

type CategoryController struct {
	BaseController
}

func (Cate *CategoryController) List(c *gin.Context) {
	var CategoryListDto dto.CategoryListDto
	if Cate.BindAndValidate(c, &CategoryListDto) {
		data, total := CategoryService.List(CategoryListDto)
		resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      CategoryListDto.Page,
			"page_size": CategoryListDto.Limit,
		})
	}
}
func (Cate *CategoryController) Create(c *gin.Context) {
	var Dto dto.CategoryCreateDto
	if Cate.BindAndValidate(c, &Dto) {
		created, _ := CategoryService.Create(Dto)
		if created.Id <= 0 {
			fail(c, ErrAddFail)
			return
		}
		resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (Cate *CategoryController) Edit(c *gin.Context) {
	var Dto dto.CategoryEditDto
	if Cate.BindAndValidate(c, &Dto) {
		affected := CategoryService.Update(Dto)
		if affected > 0 {
			ok(c, "更新成功")
			return
		}
		fail(c, ErrEditFail)
		return
	}
}

func (Cate *CategoryController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if Cate.BindAndValidate(c, &Dto) {
		data := CategoryService.InfoOfId(Dto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (Cate *CategoryController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if Cate.BindAndValidate(c, &Dto) {
		affected := CategoryService.Delete(Dto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "删除成功")
	}
}
