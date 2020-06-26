package dto

import (
	"gorm.io/datatypes"
	"time"
)

var ComicChaptersListSearchMapping = map[string]string{
	"comic_id": "comic_id",
}

type ComicChaptersListDto struct {
	Id          int                  `form:"id" json:"id"`
	Comic_id    int                  `form:"comic_id" json:"comic_id"`
	Chapter_no  int                  `form:"chapter_no" json:"chapter_no"`
	Title       string               `form:"title" json:"title"`
	Desc        datatypes.JSON `gorm:"type:json" form:"desc" json:"desc"`
	Cover       string               `form:"cover" json:"cover"`
	Upload_type int                  `form:"upload_type" json:"upload_type"`
	Is_vip      int                  `form:"is_vip" json:"is_vip"`
	CreatedAt   time.Time            `form:"created_at" json:"created_at"`
	UpdatedAt   time.Time            `form:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time           `form:"deleted_at" json:"deleted_at"`
	Page        string               `form:"page" json:"page" `
	Skip        int                  `form:"skip,default=0" json:"skip"`
	Limit       int                  `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type ComicChaptersCreateDto struct {
	Comic_id    int                  `form:"comic_id" json:"comic_id"`
	Chapter_no  int                  `form:"chapter_no" json:"chapter_no"`
	Title       string               `form:"title" json:"title"`
	Desc        datatypes.JSON `gorm:"type:json" form:"desc" json:"desc"`
	Cover       string               `form:"cover" json:"cover"`
	Upload_type int                  `form:"upload_type" json:"upload_type"`
	Is_vip      int                  `form:"is_vip" json:"is_vip"`
	CreatedAt   time.Time            `form:"created_at" json:"created_at"`
	UpdatedAt   time.Time            `form:"updated_at" json:"updated_at"`
}

type ComicChaptersEditDto struct {
	Id          int                  `uri:"id" json:"id" binding:"required"`
	Comic_id    int                  `form:"comic_id" json:"comic_id"`
	Chapter_no  int                  `form:"chapter_no" json:"chapter_no"`
	Title       string               `form:"title" json:"title"`
	Desc        datatypes.JSON `gorm:"type:json" form:"desc" json:"desc"`
	Cover       string               `form:"cover" json:"cover"`
	Upload_type int                  `form:"upload_type" json:"upload_type"`
	Is_vip      int                  `form:"is_vip" json:"is_vip"`
	UpdatedAt   time.Time            `form:"updated_at" json:"updated_at" sql:"-"`
}

type ComicChaptersUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}
