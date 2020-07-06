package model

import (
	"gorm.io/datatypes"
	"time"
)

type DistributorLevel struct {
	Id                     int            `gorm:"primary_key" json:"id"`  //分类ID
	Name                   string         `json:"name"`                   //等级名称
	Recommendtype          int            `json:"recommendtype"`          //返佣类型 1按比例返佣 2 按固定返佣
	Buyagain_switch        int            `json:"buyagain_switch"`        //复购返佣
	Auto_upgrade           int            `json:"auto_upgrade"`           //自动升级 0 否 1是
	Upgrade_conditions     datatypes.JSON `json:"upgrade_conditions"`     //升级条件
	Adaptive_degradation   int            `json:"adaptive_degradation"`   //自动降级
	Degradation_conditions datatypes.JSON `json:"degradation_conditions"` //降级条件
	Weight                 int            `json:"weight"`                 //权重
	CreatedAt              time.Time      `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt              time.Time      `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt              *time.Time     `json:"deleted_at"`
}

func (C *DistributorLevel) TableName() string {
	return "distributor_level"
}

func init() {
	//db.Register(&DistributorLevel{})
}
