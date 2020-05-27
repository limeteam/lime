package model

import (
	"lime/pkg/common/db"
	"time"
)

type Comments struct {
	Id        int        `gorm:"primary_key" json:"id"` //分类ID
	Novel_id  int        `json:"novel_id"`              //小说ID
	Username  string     `json:"username"`              //章节标题
	Content   string     `json:"content"`               //内容
	Likes     int        `json:"likes"`                 //点赞数
	Source    string     `json:"source"`                //来源
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (C *Comments) TableName() string {
	return "novel_comment"
}

func init() {
	db.Register(&Comments{})
}
