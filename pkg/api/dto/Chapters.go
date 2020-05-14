package dto

import "time"

type ChaptersListDto struct {
	Id         int        `form:"id" json:"id"`
	Novel_id   int        `form:"novel_id" json:"novel_id"`
	Chapter_no int        `form:"chapter_no" json:"chapter_no"`
	Title      string     `form:"title" json:"title"`
	Desc       string     `form:"desc" json:"desc"`
	Link       string     `form:"link" json:"link"`
	Is_vip     int        `form:"is_vip" json:"is_vip"`
	Source     string     `form:"source" json:"source"`
	Views      int        `form:"views" json:"views"`
	Text_num   int        `form:"text_num" json:"text_num"`
	Status     int        `form:"status" json:"status"`
	Try_views  int        `form:"try_views" json:"try_views"`
	CreatedAt  time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `form:"deleted_at" json:"deleted_at"`
	Page       string     `form:"page" json:"page" `
	Skip       int        `form:"skip,default=0" json:"skip"`
	Limit      int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type ChaptersCreateDto struct {
	Name                      string    `form:"name" json:"name"`
	Novel_id   int        `form:"novel_id" json:"novel_id"`
	Chapter_no int        `form:"chapter_no" json:"chapter_no"`
	Title      string     `form:"title" json:"title"`
	Desc       string     `form:"desc" json:"desc"`
	Link       string     `form:"link" json:"link"`
	Is_vip     int        `form:"is_vip" json:"is_vip"`
	Source     string     `form:"source" json:"source"`
	Views      int        `form:"views" json:"views"`
	Text_num   int        `form:"text_num" json:"text_num"`
	Status     int        `form:"status" json:"status"`
	Try_views  int        `form:"try_views" json:"try_views"`
	Is_sensitivity            int       `form:"is_sensitivity" json:"is_sensitivity"`
	CreatedAt                 time.Time `form:"created_at" json:"created_at"`
	UpdatedAt                 time.Time `form:"updated_at" json:"updated_at"`
}

type ChaptersEditDto struct {
	Id                        int       `uri:"id" json:"id" binding:"required"`
	Novel_id   int        `form:"novel_id" json:"novel_id"`
	Chapter_no int        `form:"chapter_no" json:"chapter_no"`
	Title      string     `form:"title" json:"title"`
	Desc       string     `form:"desc" json:"desc"`
	Link       string     `form:"link" json:"link"`
	Is_vip     int        `form:"is_vip" json:"is_vip"`
	Source     string     `form:"source" json:"source"`
	Views      int        `form:"views" json:"views"`
	Text_num   int        `form:"text_num" json:"text_num"`
	Status     int        `form:"status" json:"status"`
	Try_views  int        `form:"try_views" json:"try_views"`
	UpdatedAt                 time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}

type ChaptersUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}
