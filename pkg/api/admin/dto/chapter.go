package dto

import "time"

type ChapterCreateDto struct {
	Id            int       `json:"id"`
	Book_id       int       `form:"book_id" json:"book_id"`
	Title         string    `form:"title" json:"title"`
	Content       string    `form:"content" json:"content"`
	Status        int       `form:"status" json:"status"`
	Url           string    `form:"url" json:"url"`
	CreateTime    time.Time `type(datetime)" json:"create_time"`
	LastLoginTime time.Time `type(datetime)" json:"-"`
}

type ChapterEditDto struct {
	Id      int    `json:"id"`
	Book_id int    `form:"book_id" json:"book_id"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
	Status  int    `form:"status" json:"status"`
	Url     string `form:"url" json:"url"`
}
