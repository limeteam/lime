package dto

import (
	"time"
)

var ComicListSearchMapping = map[string]string{
	"channel_id":     "channel_id",
	"cid":            "category_id",
	"status":         "status",
	"online_status":  "online_status",
	"is_sensitivity": "is_sensitivity",
	"flag":           "flag",
	"name":           "name",
	//"sort": "sort",
}

type ComicsListDto struct {
	Id                 int        `form:"id" json:"id"`
	Name               string     `form:"name" json:"name" binding:"required"`
	Old_name           string     `form:"old_name" json:"old_name"`
	Channel_id         int        `form:"channel_id" json:"channel_id"`
	Category_ids       string     `form:"category_ids" json:"category_ids"`
	Desc               string     `form:"desc" json:"desc"`
	Horizontal_cover   string     `form:"horizontal_cover" json:"horizontal_cover"`
	Vertical_cover     string     `form:"vertical_cover" json:"vertical_cover"`
	Author             string     `form:"author" json:"author"`
	Source             string     `form:"source" json:"source"`
	Status             int        `form:"status" json:"status"`
	Flag               string     `form:"flag" json:"flag"`
	Chapter_price      int        `form:"chapter_price" json:"chapter_price"`
	Free_chapter       int        `form:"free_chapter" json:"free_chapter"`
	Chapter_updated_at *time.Time `form:"chapter_updated_at" json:"chapter_updated_at"`
	Views              int        `form:"views" json:"views"`
	Collect_num        int        `form:"collect_num" json:"collect_num"`
	Online_status      int        `form:"online_status" json:"online_status"`
	Is_sensitivity     int        `form:"is_sensitivity" json:"is_sensitivity"`
	CreatedAt          time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt          time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt          *time.Time `form:"deleted_at" json:"deleted_at"`
	Page               string     `form:"page" json:"page" `
	Skip               int        `form:"skip,default=0" json:"skip"`
	Limit              int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type ComicsCreateDto struct {
	Name               string     `form:"name" json:"name" binding:"required"`
	Old_name           string     `form:"old_name" json:"old_name"`
	Channel_id         int        `form:"channel_id" json:"channel_id"`
	Category_ids       string     `form:"category_ids" json:"category_ids"`
	Desc               string     `form:"desc" json:"desc"`
	Horizontal_cover   string     `form:"horizontal_cover" json:"horizontal_cover"`
	Vertical_cover     string     `form:"vertical_cover" json:"vertical_cover"`
	Author             string     `form:"author" json:"author"`
	Source             string     `form:"source" json:"source"`
	Status             int        `form:"status" json:"status"`
	Flag               string     `form:"flag" json:"flag"`
	Chapter_price      int        `form:"chapter_price" json:"chapter_price"`
	Free_chapter       int        `form:"free_chapter" json:"free_chapter"`
	Chapter_updated_at *time.Time `form:"chapter_updated_at" json:"chapter_updated_at"`
	Views              int        `form:"views" json:"views"`
	Collect_num        int        `form:"collect_num" json:"collect_num"`
	Online_status      int        `form:"online_status" json:"online_status"`
	Is_sensitivity     int        `form:"is_sensitivity" json:"is_sensitivity"`
	CreatedAt          time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt          time.Time  `form:"updated_at" json:"updated_at"`
}

type ComicsEditDto struct {
	Id                 int        `uri:"id" json:"id" binding:"required"`
	Name               string     `form:"name" json:"name" binding:"required"`
	Old_name           string     `form:"old_name" json:"old_name"`
	Channel_id         int        `form:"channel_id" json:"channel_id"`
	Category_ids       string     `form:"category_ids" json:"category_ids"`
	Desc               string     `form:"desc" json:"desc"`
	Horizontal_cover   string     `form:"horizontal_cover" json:"horizontal_cover"`
	Vertical_cover     string     `form:"vertical_cover" json:"vertical_cover"`
	Author             string     `form:"author" json:"author"`
	Source             string     `form:"source" json:"source"`
	Status             int        `form:"status" json:"status"`
	Flag               string     `form:"flag" json:"flag"`
	Chapter_price      int        `form:"chapter_price" json:"chapter_price"`
	Free_chapter       int        `form:"free_chapter" json:"free_chapter"`
	Chapter_updated_at *time.Time `form:"chapter_updated_at" json:"chapter_updated_at"`
	Views              int        `form:"views" json:"views"`
	Collect_num        int        `form:"collect_num" json:"collect_num"`
	Online_status      int        `form:"online_status" json:"online_status"`
	Is_sensitivity     int        `form:"is_sensitivity" json:"is_sensitivity"`
	UpdatedAt          time.Time  `form:"updated_at" json:"updated_at" sql:"-"`
}

type ComicsUpdateStatusDto struct {
	Id        int       `uri:"id" json:"id" binding:"required"`
	Status    int       `form:"status" json:"status"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" sql:"-"`
}
