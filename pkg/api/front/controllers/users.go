package controllers
import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/controllers"
	"lime/pkg/api/front/domain/auth"
	"lime/pkg/api/front/dto"
	"lime/pkg/api/front/service"
	"lime/pkg/api/utils/e"
)

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

func (C *UsersController) Register(c *gin.Context) {
	var Dto dto.RegisterDto
	if C.BindAndValidate(c, &Dto) {
		err := UserService.Register(Dto)
		if err != nil {
			C.Fail(c, e.ErrReg)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": true,
		})
	}
}

func (C *UsersController) Info(c *gin.Context) {
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		C.Fail(c, e.ErrUnauthorized)
		return
	}
	//lookup the metadata in redis:
	userId, err := auth.FetchAuth(metadata.TokenUuid)
	if err != nil {
		C.Fail(c, e.ErrUnauthorized)
		return
	}
	C.Resp(c, map[string]interface{}{
		"result": userId,
	})
}

func (C *UsersController) ResetPassword(c *gin.Context) {
	var Dto dto.UserEditPasswordDto
	if C.BindAndValidate(c, &Dto) {
		metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
		if err != nil {
			C.Fail(c, e.ErrUnauthorized)
			return
		}
		//lookup the metadata in redis:
		userId, err := auth.FetchAuth(metadata.TokenUuid)
		if err != nil {
			C.Fail(c, e.ErrUnauthorized)
			return
		}
		res := UserService.UpdatePassword(userId,Dto)
		C.Resp(c, map[string]interface{}{
			"data": res,
		})
	}
}
