package model

import (
	"time"
)

type Book struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Author         string    `json:"author"`
	Image          string    `json:"image"`
	Status         int       `son:"status"`
	Url            string    `json:"url"`
	Desc           string    `json:"desc"`
	CreateTime     time.Time `json:"created_time"`
	LastUpdateTime time.Time `json:"updated_time"`
}

func (d *Book) TableName() string {
	return "book"
}

func init() {
	//db.Register(&Book{})
}
