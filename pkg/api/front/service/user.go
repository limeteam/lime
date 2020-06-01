package service

import (
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
	token, _ := utils.GenerateToken(gDto.Username, gDto.Password)
	return dto.LoginInfoDto{
		Uid:               data.Id,
		Nickname:          data.Username,
		Mobile:            data.Mobile,
		UserToken:         token,
		IsVip:             data.Is_vip,
		Gender:            data.Sex,
		Avatar:            data.Faceicon,
		Remain:            "",
		GoldRemain:        "",
		SilverRemain:      "",
		CoinToday:         0,
		InviteCode:        "",
		TodayReadDuration: 0,
		IsNewUser:         0,
		CoinTipTitle:      "",
		MessageNoreadNum:  0,
		Amount:            data.Amount,
		Coin:              data.Coin,
		Level:             0,
		SignDays:          0,
		AutoSub:           "",
		BindList:          nil,
		Status:            data.Status,
		CreatedAt:         data.CreatedAt.String(),
	}, nil
}
