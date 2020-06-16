package dto

import (
	"lime/pkg/api/admin/model"
	"time"
)

var UsersListSearchMapping = map[string]string{
	"source":      "source",
	"from_market": "from_market",
	"status":      "status",
	"gender":      "gender",
	"keyword":     "keyword",
	"sort":        "sort",
}

type UsersListDto struct {
	Id            int              `form:"id" json:"id"`
	Username      string           `form:"username" json:"username"`
	Mobile        string           `form:"mobile" json:"mobile"`
	Sex           int              `form:"sex" json:"sex"`
	Password      string           `form:"password" json:"password"`
	Salt          string           `form:"salt" json:"salt"`
	Faceicon      string           `form:"faceicon" json:"faceicon"`
	Wechat        model.WechatInfo `form:"wechat" json:"wechat"`
	Email         string           `form:"email" json:"email"`
	Amount        int              `form:"amount" json:"amount"`
	Coin          int              `form:"coin" json:"coin"`
	ExemptLogin   int              `form:"exempt_login" json:"exempt_login"`
	Source        int              `form:"source" json:"source"`
	IsVip         int              `form:"is_vip" json:"is_vip"`
	ChannelId     int              `form:"channel_id" json:"channel_id"`
	Status        int              `form:"status" json:"status"`
	CreatedAt     time.Time        `form:"created_at" json:"created_at"`
	LastLoginTime time.Time        `form:"updated_at" json:"updated_at"`
	DeletedAt     *time.Time       `form:"deleted_at" json:"deleted_at"`
	Page          string           `form:"page" json:"page" `
	Skip          int              `form:"skip,default=0" json:"skip"`
	Limit         int              `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type UsersCreateDto struct {
	Username        string           `form:"username" json:"username"`
	Mobile          string           `form:"mobile" json:"mobile"`
	Sex             int              `form:"sex" json:"sex"`
	Password        string           `form:"password" json:"password"`
	Salt            string           `form:"salt" json:"salt"`
	Faceicon        string           `form:"faceicon" json:"faceicon"`
	Wechat          model.WechatInfo `gorm:"type:json form:"wechat" json:"wechat"`
	Email           string           `form:"email" json:"email"`
	Amount          int              `form:"amount" json:"amount"`
	Coin            int              `form:"coin" json:"coin"`
	Exempt_login    int              `form:"exempt_login" json:"exempt_login"`
	Source          int              `form:"source" json:"source"`
	Is_vip          int              `form:"is_vip" json:"is_vip"`
	Channel_id      int              `form:"channel_id" json:"channel_id"`
	Status          int              `form:"status" json:"status"`
	CreatedAt       time.Time        `form:"created_at" json:"created_at"`
	Last_login_time time.Time        `form:"updated_at" json:"updated_at"`
}

type UsersEditDto struct {
	Id              int              `uri:"id" json:"id" binding:"required"`
	Username        string           `form:"username" json:"username"`
	Mobile          string           `form:"mobile" json:"mobile"`
	Sex             int              `form:"sex" json:"sex"`
	Password        string           `form:"password" json:"password"`
	Salt            string           `form:"salt" json:"salt"`
	Faceicon        string           `form:"faceicon" json:"faceicon"`
	Wechat          model.WechatInfo `gorm:"type:json form:"wechat" json:"wechat"`
	Email           string           `form:"email" json:"email"`
	Amount          int              `form:"amount" json:"amount"`
	Coin            int              `form:"coin" json:"coin"`
	Exempt_login    int              `form:"exempt_login" json:"exempt_login"`
	Source          int              `form:"source" json:"source"`
	Is_vip          int              `form:"is_vip" json:"is_vip"`
	Channel_id      int              `form:"channel_id" json:"channel_id"`
	Status          int              `form:"status" json:"status"`
	CreatedAt       time.Time        `form:"created_at" json:"created_at"`
	Last_login_time time.Time        `form:"updated_at" json:"updated_at"`
}

type UsersUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}
