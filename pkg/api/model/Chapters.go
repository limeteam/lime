package model

import (
	"lime/pkg/common/db"
	"time"
)

type Chapters struct {
	Id         int        `gorm:"primary_key" json:"id"` //分类ID
	Novel_id   int        `json:"novel_id"`              //小说ID
	Chapter_no int        `json:"chapter_no"`            //章节编号
	Title      string     `json:"title"`                 //章节标题
	Desc       string     `json:"desc"`                  //章节内容
	Link       string     `json:"link"`                  //章节采集链接
	Is_vip     int        `json:"is_vip"`                //是否收费
	Source     string     `json:"source"`                //章节采集站点源
	Views      int        `json:"views"`                 //浏览次数
	Text_num   int        `json:"text_num"`              //章节字数
	Status     int        `json:"status"`                //章节采集状态0正常，1失败
	Try_views  int        `json:"try_views"`             //采集重试次数
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

func (C *Chapters) TableName() string {
	return "novel_chapter_0000"
}

func init() {
	db.Register(&Chapters{})
}
