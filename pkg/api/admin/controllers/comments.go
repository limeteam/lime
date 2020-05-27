package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var CommentsService = service.CommentService{}

type CommentsController struct {
	BaseController
}

func (C *CommentsController) List(c *gin.Context) {
	var Dto dto.CommentListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := CommentsService.List(Dto)
		C.Resp(c, map[string]interface{}{
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
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *CommentsController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := CommentsService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}
