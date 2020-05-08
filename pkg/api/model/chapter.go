package model

import (
	"lime/pkg/common/db"
	"time"
)

type Chapter struct {
	Id             int       `json:"id"`
	Book_id        int       `json:"book_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Status         int       `json:"status"`
	Url            string    `json:"url"`
	CreateTime     time.Time `json:"created_time"`
	LastUpdateTime time.Time `json:"updated_time"`
}

func (d *Chapter) TableName() string {
	return "chapter"
}

func init() {
	db.Register(&Chapter{})
}
