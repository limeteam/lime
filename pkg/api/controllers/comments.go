package controllers

import (
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"github.com/gin-gonic/gin"
)

var CommentsService = service.CommentService{}
type CommentsController struct {
	BaseController
}

func (C *CommentsController) List(c *gin.Context) {
	var Dto dto.CommentListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := CommentsService.List(Dto)
		resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      Dto.Page,
			"page_size": Dto.Limit,
		})
	}
}

func (C *CommentsController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := CommentsService.InfoOfId(Dto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *CommentsController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := CommentsService.Delete(Dto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "删除成功")
	}
}
