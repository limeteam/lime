package model

import (
	"lime/pkg/common/db"
	"time"
)

type Users struct {
	Id              int        `gorm:"primary_key" json:"id"` //分类ID
	Username        string     `json:"username"`              //用户名称
	Mobile          string     `json:"mobile"`                //手机号
	Sex             int        `json:"sex"`                   //性别
	Password        string     `json:"password"`              //密码
	Salt            string     `json:"salt"`                  //加密盐
	Faceicon        string     `json:"faceicon"`              //头像地址
	Wechat          string     `json:"wechat"`                //微信绑定信息
	Email           string     `json:"email"`                 //邮箱
	Amount          int        `json:"amount"`                //现金金额
	Coin            int        `json:"coin"`                  //金币
	Exempt_login    int        `json:"exempt_login"`          //是否免登： 0 否 1免登
	Source          int        `json:"source"`                //来源：0 PC 1 WAP 2 Android 3 IOS 4 公众号 5小程序
	Is_vip          int        `json:"is_vip"`                //是否VIP  0 否 1是
	Channel_id      int        `json:"channel_id"`            //渠道
	Status          int        `json:"status"`                //状态 1=正常 2=禁用
	CreatedAt       time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	Last_login_time time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func (U *Users) TableName() string {
	return "users"
}

func init() {
	db.Register(&Users{})
}
