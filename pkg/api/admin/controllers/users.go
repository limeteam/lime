package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var UsersService = service.UsersService{}

type AdminUsersController struct {
	BaseController
}

func (C *AdminUsersController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := UsersService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (C *AdminUsersController) Create(c *gin.Context) {
	var Dto dto.UsersCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := UsersService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *AdminUsersController) Edit(c *gin.Context) {
	var Dto dto.UsersEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := UsersService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *AdminUsersController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := UsersService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}