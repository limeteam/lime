package model

import (
	"gorm.io/datatypes"
	"time"
)

type Config struct {
	Id                int        `gorm:"primary_key" json:"id"` //分类ID
	Name              string     `json:"name"`                  //名称
	Config_code       string     `json:"config_code"`           //配置code
	Config_value datatypes.JSON     `json:"config_value"`          //配置值
	Config_group      string     `json:"config_group"`          //组
	Desc              string     `json:"desc"`                  //描述
	Status            int        `json:"status"`                //状态 -1 已经删除 0 禁用 1 启用
	CreatedAt         time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

func (C *Config) TableName() string {
	return "config"
}

func init() {
	//db.Register(&Config{})
}
