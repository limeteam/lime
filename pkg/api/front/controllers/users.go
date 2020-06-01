package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
	"lime/pkg/api/front/domain/auth"
	"lime/pkg/api/front/dto"
	"lime/pkg/api/front/service"
	"lime/pkg/api/utils/e"
)

type Authenticate struct {
	//us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}


type UsersController struct {
	controllers.BaseController
}


var UserService = service.UserService{}

func (C *UsersController) Login(c *gin.Context) {
	var Dto dto.LoginDto
	if C.BindAndValidate(c, &Dto) {
		data, err := UserService.Login(Dto)
		if err != nil {
			C.Fail(c, e.ErrLogin)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

/**
 注册用户
 */
func (C *UsersController) Register(c *gin.Context) {
	var Dto dto.LoginDto
	if C.BindAndValidate(c, &Dto) {
		data, err := UserService.Login(Dto)
		if err != nil {
			C.Fail(c, e.ErrLogin)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}
