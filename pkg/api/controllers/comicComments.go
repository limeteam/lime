package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
)

var ComicCommentsService = service.ComicCommentervice{}

type ComicCommentsController struct {
	BaseController
}

func (C *ComicCommentsController) List(c *gin.Context) {
	var Dto dto.ComicCommentListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := ComicCommentsService.List(Dto)
		resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      Dto.Page,
			"page_size": Dto.Limit,
		})
	}
}

func (C *ComicCommentsController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := ComicCommentsService.InfoOfId(Dto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *ComicCommentsController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := ComicCommentsService.Delete(Dto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "删除成功")
	}
}
