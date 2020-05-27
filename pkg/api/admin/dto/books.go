package dto

import (
	"time"
)

var BookListSearchMapping = map[string]string{
	"channel_id":     "channel_id",
	"cid":            "category_id",
	"status":         "status",
	"online_status":  "online_status",
	"is_sensitivity": "is_sensitivity",
	"flag":           "flag",
	"word":           "text_num",
	"name":           "name",
	//"sort": "sort",
}

type BooksListDto struct {
	Id                        int        `form:"id" json:"id"`
	Name                      string     `form:"name" json:"name"`
	Old_name                  string     `form:"old_name" json:"old_name"`
	Channel_id                int        `form:"channel_id" json:"channel_id"`
	Category_id               int        `form:"category_id" json:"category_id"`
	Desc                      string     `form:"desc" json:"desc"`
	Cover                     string     `form:"cover" json:"cover"`
	Author                    string     `form:"author" json:"author"`
	Source                    string     `form:"source" json:"source"`
	Split_rule                int        `form:"split_rule" json:"split_rule"`
	Upload_file               string     `form:"upload_file" json:"upload_file"`
	Status                    int        `form:"status" json:"status"`
	Flag                      string     `form:"flag" json:"flag"`
	Chapter_price             int        `form:"chapter_price" json:"chapter_price"`
	Thousand_characters_price int        `form:"thousand_characters_price" json:"thousand_characters_price"`
	Score                     int        `form:"score" json:"score"`
	Text_num                  int        `form:"text_num" json:"text_num"`
	Chapter_num               int        `form:"chapter_num" json:"chapter_num"`
	Chapter_updated_at        *time.Time `form:"chapter_updated_at" json:"chapter_updated_at"`
	Chapter_id                int        `form:"chapter_id" json:"chapter_id"`
	Chapter_title             string     `form:"chapter_title" json:"chapter_title"`
	Views                     int        `form:"views" json:"views"`
	Collect_num               int        `form:"collect_num" json:"collect_num"`
	Online_status             int        `form:"online_status" json:"online_status"`
	Is_sensitivity            int        `form:"is_sensitivity" json:"is_sensitivity"`
	CreatedAt                 time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt                 time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt                 *time.Time `form:"deleted_at" json:"deleted_at"`
	Page                      string     `form:"page" json:"page" `
	Skip                      int        `form:"skip,default=0" json:"skip"`
	Limit                     int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type BooksCreateDto struct {
	Name                      string     `form:"name" json:"name"`
	Old_name                  string     `form:"old_name" json:"old_name"`
	Channel_id                int        `form:"channel_id" json:"channel_id"`
	Category_id               int        `form:"category_id" json:"category_id"`
	Desc                      string     `form:"desc" json:"desc"`
	Cover                     string     `form:"cover" json:"cover"`
	Author                    string     `form:"author" json:"author"`
	Source                    string     `form:"source" json:"source"`
	Split_rule                int        `form:"split_rule" json:"split_rule"`
	Upload_file               string     `form:"upload_file" json:"upload_file"`
	Status                    int        `form:"status" json:"status"`
	Flag                      string     `form:"flag" json:"flag"`
	Chapter_price             int        `form:"chapter_price" json:"chapter_price"`
	Thousand_characters_price int        `form:"thousand_characters_price" json:"thousand_characters_price"`
	Score                     int        `form:"score" json:"score"`
	Text_num                  int        `form:"text_num" json:"text_num"`
	Chapter_num               int        `form:"chapter_num" json:"chapter_num"`
	Chapter_updated_at        *time.Time `form:"chapter_updated_at" json:"chapter_updated_at"`
	Chapter_id                int        `form:"chapter_id" json:"chapter_id"`
	Chapter_title             string     `form:"chapter_title" json:"chapter_title"`
	Views                     int        `form:"views" json:"views"`
	Collect_num               int        `form:"collect_num" json:"collect_num"`
	Online_status             int        `form:"online_status" json:"online_status"`
	Is_sensitivity            int        `form:"is_sensitivity" json:"is_sensitivity"`
	CreatedAt                 time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt                 time.Time  `form:"updated_at" json:"updated_at"`
}

type BooksEditDto struct {
	Id                        int        `uri:"id" json:"id" binding:"required"`
	Name                      string     `form:"name" json:"name" binding:"required"`
	Old_name                  string     `form:"old_name" json:"old_name"`
	Channel_id                int        `form:"channel_id" json:"channel_id"`
	Category_id               int        `form:"category_id" json:"category_id"`
	Desc                      string     `form:"desc" json:"desc"`
	Cover                     string     `form:"cover" json:"cover"`
	Author                    string     `form:"author" json:"author"`
	Source                    string     `form:"source" json:"source"`
	Split_rule                int        `form:"split_rule" json:"split_rule"`
	Upload_file               string     `form:"upload_file" json:"upload_file"`
	Status                    int        `form:"status" json:"status"`
	Flag                      string     `form:"flag" json:"flag"`
	Chapter_price             int        `form:"chapter_price" json:"chapter_price"`
	Thousand_characters_price int        `form:"thousand_characters_price" json:"thousand_characters_price"`
	Score                     int        `form:"score" json:"score"`
	Text_num                  int        `form:"text_num" json:"text_num"`
	Chapter_num               int        `form:"chapter_num" json:"chapter_num"`
	Chapter_updated_at        *time.Time `form:"chapter_updated_at" json:"chapter_updated_at"`
	Chapter_id                int        `form:"chapter_id" json:"chapter_id"`
	Chapter_title             string     `form:"chapter_title" json:"chapter_title"`
	Views                     int        `form:"views" json:"views"`
	Collect_num               int        `form:"collect_num" json:"collect_num"`
	Online_status             int        `form:"online_status" json:"online_status"`
	Is_sensitivity            int        `form:"is_sensitivity" json:"is_sensitivity"`
	UpdatedAt                 time.Time  `form:"updated_at" json:"updated_at" sql:"-"`
}

type BooksUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}