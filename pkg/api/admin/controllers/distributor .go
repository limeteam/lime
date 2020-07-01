package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var DistributorService = service.DistributorService{}

type DistributorController struct {
	BaseController
}

func (C *DistributorController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if C.BindAndValidate(c, &Dto) {
		data, total := DistributorService.List(Dto)
		C.Resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (C *DistributorController) Create(c *gin.Context) {
	var Dto dto.DistributorCreateDto
	if C.BindAndValidate(c, &Dto) {
		created, _ := DistributorService.Create(Dto)
		if created.Id <= 0 {
			C.Fail(c, e.ErrAddFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (C *DistributorController) Edit(c *gin.Context) {
	var Dto dto.DistributorEditDto
	if C.BindAndValidate(c, &Dto) {
		affected := DistributorService.Update(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *DistributorController) UpdateStatus(c *gin.Context) {
	var Dto dto.DistributorUpdateStatusDto
	if C.BindAndValidate(c, &Dto) {
		affected := DistributorService.UpdateStatus(Dto)
		if affected > 0 {
			C.Ok(c, "更新成功")
			return
		}
		C.Fail(c, e.ErrEditFail)
		return
	}
}

func (C *DistributorController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if C.BindAndValidate(c, &Dto) {
		data := DistributorService.InfoOfId(Dto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *DistributorController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if C.BindAndValidate(c, &Dto) {
		affected := DistributorService.Delete(Dto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Ok(c, "删除成功")
	}
}
