package model

import (
	"time"
)

type Distributor struct {
	Id                   int        `gorm:"primary_key" json:"id"` //分类ID
	Recommender_id       int        `json:"recommender_id"`        //推荐人ID
	Distributor_id       int        `json:"distributor_id"`        //分销商ID
	Name                 string     `json:"name"`                  //姓名
	Mobile               string     `json:"mobile"`                //手机
	Distribution_level   int        `json:"distribution_level"`    //分销等级
	Commission_withdrawn int        `json:"commission_withdrawn"`  //已经提现佣金
	Commission_available int        `json:"commission_available"`  //可提现佣金
	Status               int        `json:"status"`                //状态：1 已审核 2 待审核 3 已拒绝
	CreatedAt            time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt            *time.Time `json:"deleted_at"`
}

func (C *Distributor) TableName() string {
	return "distributor"
}

func init() {
	//db.Register(&Distributor{})
}
