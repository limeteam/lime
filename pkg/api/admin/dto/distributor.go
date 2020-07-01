package dto

import (
	"time"
)

var DistributorSearchMapping = map[string]string{
	"name": "name",
}

type DistributorListDto struct {
	Id                   int        `form:"id" json:"id"`
	Recommender_id       int        `form:"recommender_id" json:"recommender_id"`             //推荐人ID
	Distributor_id       int        `form:"distributor_id" json:"distributor_id"`             //分销商ID
	Name                 string     `form:"name" json:"name"`                                 //姓名
	Mobile               string     `form:"mobile" json:"mobile"`                             //手机
	Distribution_level   int        `form:"distribution_level" json:"distribution_level"`     //分销等级
	Commission_withdrawn int        `form:"commission_withdrawn" json:"commission_withdrawn"` //已经提现佣金
	Commission_available int        `form:"commission_available" json:"commission_available"` //可提现佣金
	Status               int        `form:"status" json:"status"`                             //状态：1 已审核 2 待审核 3 已拒绝
	CreatedAt            time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt            time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt            *time.Time `form:"deleted_at" json:"deleted_at"`
	Page                 string     `form:"page" json:"page" `
	Skip                 int        `form:"skip,default=0" json:"skip"`
	Limit                int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type DistributorCreateDto struct {
	Recommender_id       int       `form:"recommender_id" json:"recommender_id"`             //推荐人ID
	Distributor_id       int       `form:"distributor_id" json:"distributor_id"`             //分销商ID
	Name                 string    `form:"name" json:"name"`                                 //姓名
	Mobile               string    `form:"mobile" json:"mobile"`                             //手机
	Distribution_level   int       `form:"distribution_level" json:"distribution_level"`     //分销等级
	Commission_withdrawn int       `form:"commission_withdrawn" json:"commission_withdrawn"` //已经提现佣金
	Commission_available int       `form:"commission_available" json:"commission_available"` //可提现佣金
	Status               int       `form:"status" json:"status"`                             //状态：1 已审核 2 待审核 3 已拒绝
	CreatedAt            time.Time `form:"created_at" json:"created_at"`
	UpdatedAt            time.Time `form:"updated_at" json:"updated_at"`
}

type DistributorEditDto struct {
	Id                   int       `uri:"id" json:"id" binding:"required"`
	Recommender_id       int       `form:"recommender_id" json:"recommender_id"`             //推荐人ID
	Distributor_id       int       `form:"distributor_id" json:"distributor_id"`             //分销商ID
	Name                 string    `form:"name" json:"name"`                                 //姓名
	Mobile               string    `form:"mobile" json:"mobile"`                             //手机
	Distribution_level   int       `form:"distribution_level" json:"distribution_level"`     //分销等级
	Commission_withdrawn int       `form:"commission_withdrawn" json:"commission_withdrawn"` //已经提现佣金
	Commission_available int       `form:"commission_available" json:"commission_available"` //可提现佣金
	Status               int       `form:"status" json:"status"`                             //状态：1 已审核 2 待审核 3 已拒绝
	UpdatedAt            time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}

type DistributorUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}
