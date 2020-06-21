package dto

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type LoginDto struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type RegisterDto struct {
	Username    string `json:"username"`
	Email		string `json:"email"`
	Mobile		string `json:"mobile"`
	Password    string `json:"password"`
}

type LoginInfoDto struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires int `json:"at_expires"`
	RtExpires int `json:"rt_expires"`
}

type InfoDto struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Faceicon string `json:"faceicon"`
}

type BindList struct {
	Label   string `json:"label"`
	Action  string `json:"action"`
	Status  int    `json:"status"`
	Display string `json:"display"`
}

type UserEditPasswordDto struct {
	Password string `form:"new_password" json:"password" binding:"required,pwdValidate"`
	//RePassword string `form:"re_password" json:"re_password" binding:"required,pwdValidate"`
}

func pwdValidate(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*]{6,}$`)
	val := fl.Field().String()
	if !reg.Match([]byte(val)) {
		return false
	}
	return true
}
