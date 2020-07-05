package dto

import (
	"gorm.io/datatypes"
	"time"
)

var ConfigSearchMapping = map[string]string{
	"name": "name",
}

type ConfigListDto struct {
	Id           int            `form:"id" json:"id"`
	Name         string         `form:"name" json:"name"`                 //名称
	Config_code  string         `form:"config_code" json:"config_code"`   //配置code
	Config_value datatypes.JSON `form:"config_value" json:"config_value"` //配置值
	Config_group string         `form:"config_group" json:"config_group"` //组
	Desc         string         `form:"desc" json:"desc"`                 //描述
	Status       int            `form:"status"  json:"status"`            //状态 -1 已经删除 0 禁用 1 启用
	CreatedAt    time.Time      `form:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `form:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time     `form:"deleted_at" json:"deleted_at"`
	Page         string         `form:"page" json:"page" `
	Skip         int            `form:"skip,default=0" json:"skip"`
	Limit        int            `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type ConfigCreateDto struct {
	Name         string         `form:"name" json:"name"`                 //名称
	Config_code  string         `form:"config_code" json:"config_code"`   //配置code
	Config_value datatypes.JSON `form:"config_value" json:"config_value"` //配置值
	Config_group string         `form:"config_group" json:"config_group"` //组
	Desc         string         `form:"desc" json:"desc"`                 //描述
	Status       int            `form:"status"  json:"status"`            //状态
	CreatedAt    time.Time      `form:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `form:"updated_at" json:"updated_at"`
}

type ConfigEditDto struct {
	Id           int            `uri:"id" json:"id" binding:"required"`
	Name         string         `form:"name" json:"name"`                 //名称
	Config_code  string         `form:"config_code" json:"config_code"`   //配置code
	Config_value datatypes.JSON `form:"config_value" json:"config_value"` //配置值
	Config_group string         `form:"config_group" json:"config_group"` //组
	Desc         string         `form:"desc" json:"desc"`                 //描述
	Status       int            `form:"status"  json:"status"`            //状态
	UpdatedAt    time.Time      `form:"updated_at" json:"updated_at" sql:"-"`
}

type ConfigUpdateDto struct {
	Config_code  string         `uri:"config_code" form:"config_code" json:"config_code"`   //配置code
	Config_value datatypes.JSON `form:"config_value" json:"config_value"` //配置值
}

type ConfigUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}

type ConfigGetByCodeDto struct {
	Code string `uri:"code" json:"code" binding:"required"`
}
