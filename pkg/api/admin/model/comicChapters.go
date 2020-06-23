package model

import (
	"time"
)

type ComicChapters struct {
	Id          int        `gorm:"primary_key" json:"id"` //分类ID
	Comic_id    int        `json:"comic_id"`              //漫画ID
	Chapter_no  int        `json:"chapter_no"`            //章节编号
	Title       string     `json:"title"`                 //章节标题
	Desc        string     `json:"desc"`                  //章节图
	Cover       string     `json:"cover"`                 //封面
	Upload_type int        `json:"upload_type"`           //上传类型 1多图上传 2链接上传
	Is_vip      int        `json:"is_vip"`                //是否收费
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (C *ComicChapters) TableName() string {
	return "comic_chapter"
}

func init() {
	//db.Register(&ComicChapters{})
}
