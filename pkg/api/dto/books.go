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

type BookInfoTag struct {
	Tab   string `json:"tab"`
	Color string `json:"color"`
}

type BookAttribute struct {
	Popularity      string `json:"popularity"`
	PopularityTitle string `json:"popularity_title"`
	Reading         string `json:"reading"`
	ReadingTitle    string `json:"reading_title"`
	Score           string `json:"score"`
	ScoreTitle      string `json:"score_title"`
}

type BookComment struct {
	CommentId int        `json:"comment_id"`
	Uid       int        `json:"uid"`
	Nickname  string     `json:"nickname"`
	Avatar    string     `json:"avatar"`
	Time      *time.Time `json:"time"`
	LikeNum   int        `json:"like_num"`
	Content   string     `json:"content"`
	ReplyInfo string     `json:"reply_info"`
	IsVip     int        `json:"is_vip"`
}

type BookInfoAllLookInfos struct {
	BookId      int           `json:"book_id"`
	Name        string        `json:"name"`
	Cover       string        `json:"cover"`
	Author      string        `json:"author"`
	Description string        `json:"description"`
	Views       int           `json:"views"`
	Tag         []BookInfoTag `json:"tag"`
	Finished    string        `json:"finished"`
	Flag        string        `json:"flag"`
	TotalWords  string        `json:"total_words"`
	IsVip       int           `json:"is_vip"`
	IsBaoyue    int           `json:"is_baoyue"`
	IsFinished  int           `json:"is_finished"`
}

type BookInfoLabels struct {
	RecommendId int                    `json:"recommend_id"`
	Label       string                 `json:"label"`
	Style       int                    `json:"style"`
	CanMore     bool                   `json:"can_more"`
	CanRefresh  bool                   `json:"can_refresh"`
	Total       int                    `json:"total"`
	List        []BookInfoAllLookInfos `json:"list"`
}
type BookInfoDto struct {
	BookId          int            `json:"book_id"`
	Name            string         `json:"name"`
	Cover           string         `json:"cover"`
	Author          string         `json:"author"`
	Description     string         `json:"description"`
	DisplayLabel    string         `json:"display_label"`
	Tags            []BookInfoTag  `json:"tag",omitempty`
	Finished        string         `json:"finished"`
	Flag            string         `json:"flag"`
	TotalWords      string         `json:"total_words"`
	TotalComment    string         `json:"total_comment"`
	ChapterLabel    string         `json:"chapter_label"`
	LastChapterTime string         `json:"last_chapter_time"`
	LastChapter     string         `json:"last_chapter"`
	IsFinished      int            `json:"is_finished"`
	Score           int         `json:"score"`
	Attribute       BookAttribute  `json:"attribute"`
	Comment         []BookComment  `json:"comment",omitempty`
	Labels          BookInfoLabels `json:"labels"`
	IsCollect       int            `json:"is_collect"`
}
