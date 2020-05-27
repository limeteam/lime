package model

import (
	"lime/pkg/common/db"
	"time"
)

type ComicComments struct {
	Id        int        `gorm:"primary_key" json:"id"` //分类ID
	Comic_id  int        `json:"comic_id"`              //漫画ID
	Username  string     `json:"username"`              //章节标题
	Content   string     `json:"content"`               //内容
	Source    string     `json:"source"`                //来源
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (C *ComicComments) TableName() string {
	return "comic_comment"
}

func init() {
	db.Register(&ComicComments{})
}
