package dto

import (
	"time"
)

type BookInfoTag struct {
	Tab   string `json:"tab"`
	Color string `json:"color"`
}

var BookListSearchMapping = map[string]string{
	"type":     "category_id",
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
