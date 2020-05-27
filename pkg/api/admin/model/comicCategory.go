package model

import (
	"lime/pkg/common/db"
	"time"
)

type ComicCategory struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Comic_num int        `json:"comic_num"`
	Sort      int        `json:"sort"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (C *ComicCategory) TableName() string {
	return "comic_category"
}

func init() {
	db.Register(&ComicCategory{})
}
