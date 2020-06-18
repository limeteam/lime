package service

import (
	"errors"
	"lime/pkg/api/front/dao"
	"lime/pkg/api/front/dto"
	"lime/pkg/api/utils"
)

type UserService struct{}

var UserDao = dao.UsersDao{}

func (U UserService) Login(gDto dto.LoginDto) (loginInfo dto.LoginInfoDto, err error) {
	data := UserDao.GetUserByUsername(gDto.Username)
	if data.Password != utils.EncodeMD5(gDto.Password+data.Salt) {
		return loginInfo, err
	}
	//token, _ := utils.GenerateToken(gDto.Username, gDto.Password)
	return dto.LoginInfoDto{

	}, nil
}

func (U UserService) Register(gDto dto.RegisterDto) (err error) {
	data := UserDao.GetUserByUsername(gDto.Username)
	if data.Username != "" {
		return errors.New("已经有")
	}
	return  nil
}


