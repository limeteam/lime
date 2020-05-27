package dto

import "time"

type BookCreateDto struct {
	Name          string    `form:"name" json:"name"`
	Author        string    `form:"author" json:"author"`
	Image         string    `form:"image" json:"image"`
	Status        int       `form:"status" json:"status"`
	Url           string    `form:"url" json:"url"`
	Desc          string    `form:"desc" json:"desc"`
	CreateTime    time.Time `type(datetime)" json:"create_time"`
	LastLoginTime time.Time `type(datetime)" json:"-"`
}

type BookListDto struct {
	Name   string `form:"name" json:"name"`
	Author string `form:"author" json:"author"`
	Image  string `form:"image" json:"image"`
	Status int    `form:"status" json:"status"`
	Url    string `form:"url" json:"url"`
	Page   string `form:"page" json:"page" `
	Skip   int    `form:"skip,default=0" json:"skip"`
	Limit  int    `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type BookEditDto struct {
	Id            int       `json:"id"`
	Name          string    `form:"name" json:"name"`
	Author        string    `form:"author" json:"author"`
	Image         string    `form:"image" json:"image"`
	Status        int       `form:"status" json:"status"`
	Url           string    `form:"url" json:"url"`
	Desc          string    `form:"desc" json:"desc"`
	CreateTime    time.Time `type(datetime)" json:"create_time"`
	LastLoginTime time.Time `type(datetime)" json:"-"`
}
