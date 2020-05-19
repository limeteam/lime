package dto

import "time"

type ComicCategoryListDto struct {
	Name      string     `form:"name" json:"name" `
	Comic_num  int        `form:"comic_num,default=0" json:"comic_num" `
	Sort      int        `form:"sort" json:"sort" `
	CreatedAt time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `form:"deleted_at" json:"deleted_at"`
	Page      string     `form:"page" json:"page" `
	Skip      int        `form:"skip,default=0" json:"skip"`
	Limit     int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type ComicCategoryCreateDto struct {
	Name      string    `form:"name" json:"name" `
	Comic_num  int        `form:"comic_num,default=0" json:"comic_num" `
	Sort      int       `form:"sort" json:"sort" `
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
}

type ComicCategoryEditDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Name      string    `form:"name" json:"name" binding:"required"`
	Comic_num  int        `form:"comic_num,default=0" json:"comic_num" `
	Sort      int       `form:"sort" json:"sort" `
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}
