package model

import (
	"lime/pkg/common/db"
	"time"
)

type Category struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	ChannelId string    `json:"channel_id"`
	NovelNum  string    `json:"novel_num"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (C *Category) TableName() string {
	return "novel_category"
}

func init() {
	db.Register(&Category{})
}
