package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var DistributorLevelService = service.DistributorLevelService{}

type DistributorLevelController struct {
	BaseController
}

func (C *DistributorLevelController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := DistributorLevelService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (C *DistributorLevelController) Create(c *gin.Context) {
	var Dto dto.DistributorLevelCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := DistributorLevelService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *DistributorLevelController) Edit(c *gin.Context) {
	var Dto dto.DistributorLevelEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := DistributorLevelService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *DistributorLevelController) UpdateStatus(c *gin.Context) {
	var Dto dto.DistributorLevelUpdateStatusDto
	if C.BindAndValidate(c, &Dto) {
		affected := DistributorLevelService.UpdateStatus(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *DistributorLevelController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := DistributorLevelService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *DistributorLevelController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := DistributorLevelService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}