package service

import (
	"errors"
	"lime/pkg/api/front/dao"
	"lime/pkg/api/front/domain/auth"
	"lime/pkg/api/front/dto"
)

type UserService struct{
	tk auth.TokenInterface
	rd auth.AuthInterface
}

var UserDao = dao.UsersDao{}

func (U UserService) Login(gDto dto.LoginDto) (loginInfo dto.LoginInfoDto, err error) {
	var tokenErr = map[string]string{}
	u, userErr := UserDao.GetUserByEmailAndPassword(gDto.Username,gDto.Password)
	if userErr != nil {
		//c.JSON(http.StatusInternalServerError, userErr)
		return loginInfo , errors.New("no user")
	}
	ts, tErr := U.tk.CreateToken(u.Id)
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		//c.JSON(http.StatusUnprocessableEntity, tErr.Error())
		return loginInfo ,tErr
	}
	saveErr := U.rd.CreateAuth(u.Id, ts)
	if saveErr != nil {
		//c.JSON(http.StatusInternalServerError, saveErr.Error())
		return loginInfo ,saveErr
	}
	//userData := make(map[string]interface{})
	//userData["access_token"] = ts.AccessToken
	//userData["refresh_token"] = ts.RefreshToken
	//userData["id"] = u.ID
	//userData["first_name"] = u.FirstName
	//userData["last_name"] = u.LastName
	return dto.LoginInfoDto{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
		AtExpires:    0,
		RtExpires:    0,
	},nil
}

func (U UserService) Register(gDto dto.RegisterDto) (err error) {
	data := UserDao.GetUserByUsername(gDto.Username)
	if data.Username != "" {
		return errors.New("已经有")
	}
	return  nil
}


