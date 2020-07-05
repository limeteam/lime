package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var ConfigService = service.ConfigService{}

type ConfigController struct {
	BaseController
}

func (C *ConfigController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := ConfigService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (C *ConfigController) Create(c *gin.Context) {
	var Dto dto.ConfigCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := ConfigService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *ConfigController) Edit(c *gin.Context) {
	var Dto dto.ConfigEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := ConfigService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ConfigController) EditByCode(c *gin.Context) {
	var Dto dto.ConfigUpdateDto
	if C.BindAndValidate(c, &Dto) {
		affected := ConfigService.UpdateByCode(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ConfigController) UpdateStatus(c *gin.Context) {
	var Dto dto.ConfigUpdateStatusDto
	if C.BindAndValidate(c, &Dto) {
		affected := ConfigService.UpdateStatus(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *ConfigController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := ConfigService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *ConfigController) GetByCode(c *gin.Context) {
	var Dto dto.ConfigGetByCodeDto
	if C.BindAndValidate(c, &Dto) {
		data := ConfigService.InfoOfCode(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data.Config_value,
		})
	}
}

func (C *ConfigController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := ConfigService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}
