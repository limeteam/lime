package controllers
import (
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"github.com/gin-gonic/gin"
)

var CategoryService = service.CategoryService{}

type CategoryController struct {
	BaseController
}

func (C *CategoryController) List(c *gin.Context) {
	var listDto dto.CategoryListDto
	if C.BindAndValidate(c, &listDto) {
		data, total := CategoryService.List(listDto)
		resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      listDto.Page,
			"page_size": listDto.Limit,
		})
	}
}
