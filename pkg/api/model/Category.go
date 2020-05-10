package model

import (
	"lime/pkg/common/db"
	"time"
)

type Category struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ChannelId int    `json:"channel_id"`
	NovelNum  int    `json:"novel_num"`
	Sort      int    `json:"sort"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time  `json:"deleted_at"`
}

func (C *Category) TableName() string {
	return "novel_category"
}

func init() {
	db.Register(&Category{})
}
