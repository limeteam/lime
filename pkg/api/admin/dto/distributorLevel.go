package dto

import (
	"time"
)

var DistributorLevelSearchMapping = map[string]string{
	"name": "name",
}

type DistributorLevelListDto struct {
	Id                     int        `form:"id" json:"id"`
	Name                   string     `form:"name" json:"name"`                                     //等级名称
	Repurchase_commission  string     `form:"repurchase_commission" json:"repurchase_commission"`   //复购返佣
	Recommendtype          int        `form:"recommendtype" json:"recommendtype"`                   //返佣类型 1按比例返佣 2 按固定返佣
	Auto_upgrade           int        `form:"auto_upgrade" json:"auto_upgrade"`                     //自动升级 0 否 1是
	Upgrade_conditions     string     `form:"upgrade_conditions" json:"upgrade_conditions"`         //升级条件
	Adaptive_degradation   int        `form:"adaptive_degradation" json:"adaptive_degradation"`     //自动降级
	Degradation_conditions string     `form:"degradation_conditions" json:"degradation_conditions"` //降级条件
	Weight                 int        `form:"weight" json:"weight"`                                 //权重
	CreatedAt              time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt              time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt              *time.Time `form:"deleted_at" json:"deleted_at"`
	Page                   string     `form:"page" json:"page" `
	Skip                   int        `form:"skip,default=0" json:"skip"`
	Limit                  int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type DistributorLevelCreateDto struct {
	Name                   string    `form:"name" json:"name"`                                     //等级名称
	Repurchase_commission  string    `form:"repurchase_commission" json:"repurchase_commission"`   //复购返佣
	Recommendtype          int       `form:"recommendtype" json:"recommendtype"`                   //返佣类型 1按比例返佣 2 按固定返佣
	Auto_upgrade           int       `form:"auto_upgrade" json:"auto_upgrade"`                     //自动升级 0 否 1是
	Upgrade_conditions     string    `form:"upgrade_conditions" json:"upgrade_conditions"`         //升级条件
	Adaptive_degradation   int       `form:"adaptive_degradation" json:"adaptive_degradation"`     //自动降级
	Degradation_conditions string    `form:"degradation_conditions" json:"degradation_conditions"` //降级条件
	Weight                 int       `form:"weight" json:"weight"`                                 //权重
	CreatedAt              time.Time `form:"created_at" json:"created_at"`
	UpdatedAt              time.Time `form:"updated_at" json:"updated_at"`
}

type DistributorLevelEditDto struct {
	Id                     int       `uri:"id" json:"id" binding:"required"`
	Name                   string    `form:"name" json:"name"`                                     //等级名称
	Repurchase_commission  string    `form:"repurchase_commission" json:"repurchase_commission"`   //复购返佣
	Recommendtype          int       `form:"recommendtype" json:"recommendtype"`                   //返佣类型 1按比例返佣 2 按固定返佣
	Auto_upgrade           int       `form:"auto_upgrade" json:"auto_upgrade"`                     //自动升级 0 否 1是
	Upgrade_conditions     string    `form:"upgrade_conditions" json:"upgrade_conditions"`         //升级条件
	Adaptive_degradation   int       `form:"adaptive_degradation" json:"adaptive_degradation"`     //自动降级
	Degradation_conditions string    `form:"degradation_conditions" json:"degradation_conditions"` //降级条件
	Weight                 int       `form:"weight" json:"weight"`                                 //权重
	UpdatedAt              time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}

type DistributorLevelUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}
