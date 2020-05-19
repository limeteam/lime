package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
)

var ComicCategoryService = service.ComicCategoryService{}

type ComicCategoryController struct {
	BaseController
}

func (Cate *ComicCategoryController) List(c *gin.Context) {
	var Dto dto.ComicCategoryListDto
	if Cate.BindAndValidate(c, &Dto) {
		data, total := ComicCategoryService.List(Dto)
		resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      Dto.Page,
			"page_size": Dto.Limit,
		})
	}
}
func (Cate *ComicCategoryController) Create(c *gin.Context) {
	var Dto dto.ComicCategoryCreateDto
	if Cate.BindAndValidate(c, &Dto) {
		created, _ := ComicCategoryService.Create(Dto)
		if created.Id <= 0 {
			fail(c, ErrAddFail)
			return
		}
		resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (Cate *ComicCategoryController) Edit(c *gin.Context) {
	var Dto dto.ComicCategoryEditDto
	if Cate.BindAndValidate(c, &Dto) {
		affected := ComicCategoryService.Update(Dto)
		if affected > 0 {
			ok(c, "更新成功")
			return
		}
		fail(c, ErrEditFail)
		return
	}
}

func (Cate *ComicCategoryController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if Cate.BindAndValidate(c, &Dto) {
		data := ComicCategoryService.InfoOfId(Dto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (Cate *ComicCategoryController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if Cate.BindAndValidate(c, &Dto) {
		affected := ComicCategoryService.Delete(Dto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "删除成功")
	}
}
