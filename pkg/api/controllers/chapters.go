package controllers
import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
)

var ChaptersService = service.ChaptersService{}

type ChaptersController struct {
	BaseController
}

func (Chapter *ChaptersController) List(c *gin.Context) {
	var Dto dto.GeneralListDto
	if Chapter.BindAndValidate(c, &Dto) {
		data, total := ChaptersService.List(Dto)
		resp(c, map[string]interface{}{
			"result": data,
			"total":  total,
		})
	}
}
func (Chapter *ChaptersController) Create(c *gin.Context) {
	var Dto dto.ChaptersCreateDto
	if Chapter.BindAndValidate(c, &Dto) {
		created, _ := ChaptersService.Create(Dto)
		if created.Id <= 0 {
			fail(c, ErrAddFail)
			return
		}
		resp(c, map[string]interface{}{
			"id": created.Id,
		})
	}
}

func (Chapter *ChaptersController) Edit(c *gin.Context) {
	var Dto dto.ChaptersEditDto
	if Chapter.BindAndValidate(c, &Dto) {
		affected := ChaptersService.Update(Dto)
		if affected > 0 {
			ok(c, "更新成功")
			return
		}
		fail(c, ErrEditFail)
		return
	}
}

func (Chapter *ChaptersController) UpdateStatus(c *gin.Context) {
	var Dto dto.ChaptersUpdateStatusDto
	if Chapter.BindAndValidate(c, &Dto) {
		affected := ChaptersService.UpdateStatus(Dto)
		if affected > 0 {
			ok(c, "更新成功")
			return
		}
		fail(c, ErrEditFail)
		return
	}
}

func (Chapter *ChaptersController) Get(c *gin.Context) {
	var Dto dto.GeneralGetDto
	if Chapter.BindAndValidate(c, &Dto) {
		data := ChaptersService.InfoOfId(Dto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (Chapter *ChaptersController) Delete(c *gin.Context) {
	var Dto dto.GeneralDelDto
	if Chapter.BindAndValidate(c, &Dto) {
		affected := ChaptersService.Delete(Dto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		ok(c, "删除成功")
	}
}