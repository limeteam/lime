package v1

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/controllers"
	"lime/pkg/api/dto"
)

type UsersController struct {
	controllers.BaseController
}

func (U *UsersController) Login(c *gin.Context) {
	var Dto dto.GeneralListDto
	if U.BindAndValidate(c, &Dto) {
		//data, total := BooksService.List(Dto)
		//Base.Resp(c, map[string]interface{}{
		//	"result": data,
		//	//"total":  total,
		//})
	}
}
