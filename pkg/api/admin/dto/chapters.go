package dto

import "time"

var ChaptersListSearchMapping = map[string]string{
	"novel_id": "novel_id",
}

type ChaptersListDto struct {
	Id         int        `form:"id" json:"id"`
	Novel_id   int        `uri::"novel_id" json:"novel_id"`
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
	Name           string    `form:"name" json:"name"`
	Novel_id       int       `form:"novel_id" json:"novel_id"`
	Chapter_no     int       `form:"chapter_no" json:"chapter_no"`
	Title          string    `form:"title" json:"title"`
	Desc           string    `form:"desc" json:"desc"`
	Link           string    `form:"link" json:"link"`
	Is_vip         int       `form:"is_vip" json:"is_vip"`
	Source         string    `form:"source" json:"source"`
	Views          int       `form:"views" json:"views"`
	Text_num       int       `form:"text_num" json:"text_num"`
	Status         int       `form:"status" json:"status"`
	Try_views      int       `form:"try_views" json:"try_views"`
	Is_sensitivity int       `form:"is_sensitivity" json:"is_sensitivity"`
	CreatedAt      time.Time `form:"created_at" json:"created_at"`
	UpdatedAt      time.Time `form:"updated_at" json:"updated_at"`
}

type ChaptersEditDto struct {
	Id         int       `uri:"id" json:"id" binding:"required"`
	Novel_id   int       `form:"novel_id" json:"novel_id"`
	Chapter_no int       `form:"chapter_no" json:"chapter_no"`
	Title      string    `form:"title" json:"title"`
	Desc       string    `form:"desc" json:"desc"`
	Link       string    `form:"link" json:"link"`
	Is_vip     int       `form:"is_vip" json:"is_vip"`
	Source     string    `form:"source" json:"source"`
	Views      int       `form:"views" json:"views"`
	Text_num   int       `form:"text_num" json:"text_num"`
	Status     int       `form:"status" json:"status"`
	Try_views  int       `form:"try_views" json:"try_views"`
	UpdatedAt  time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}

type ChaptersUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}

type ChapterBookInfo struct {
	BookId     int    `json:"book_id"`
	Name       string `json:"name"`
	IsFinished int    `json:"is_finished"`
}

type ChapterInfoDto struct {
	Book         ChapterBookInfo `json:"book"`
	ChapterTitle string          `json:"chapter_title"`
	Content      string          `json:"content"`
	IsPreview    int             `json:"is_preview"`
	LastChapter  int             `json:"last_chapter"`
	NextChapter  string          `json:"next_chapter"`
	ChapterId    int             `json:"chapter_id"`
	UpdateTime   string          `json:"update_time"`
}

type ChaptersFrontLists struct {
	BookId       int    `json:"book_id"`
	ChapterId    int    `json:"chapter_id"`
	ChapterTitle string `json:"chapter_title"`
	Words        int    `json:"words"`
	IsVip        int    `json:"is_vip"`
	CanRead      int    `json:"can_read"`
	IsPreview    int    `json:"is_preview"`
	Tag          struct {
		Tab   string `json:"tab"`
		Color string `json:"color"`
	} `json:"tag"`
	UpdateTime   string `json:"update_time"`
	DisplayOrder int `json:"display_order"`
	LastChapter  int `json:"last_chapter"`
	NextChapter  int `json:"next_chapter"`
}
type ChaptersFrontListDto struct {
	BookId       int                  `json:"book_id"`
	Name         string               `json:"name"`
	Cover        string               `json:"cover"`
	Description  string               `json:"description"`
	TotalChapter int                  `json:"total_chapter"`
	ChapterList  []ChaptersFrontLists `json:"chapter_list"`
}
