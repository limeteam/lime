package model

import (
	"time"
)

type Comics struct {
	Id                 int        `gorm:"primary_key" json:"id"` //分类ID
	Name               string     `json:"name"`                  //漫画名
	Old_name           string     `json:"old_name"`              //原名
	Channel_id         int        `json:"channel_id"`            //所属频道
	Category_ids       string     `json:"category_ids"`          //漫画分类
	Desc               string     `json:"desc"`                  //漫画简介
	Horizontal_cover   string     `json:"horizontal_cover"`      //横版封面
	Vertical_cover     string     `json:"vertical_cover"`        //竖版封面
	Author             string     `json:"author"`                //漫画作者
	Source             string     `json:"source"`                //来源
	Status             int        `json:"status"`                //小说状态，0连载中，1已完结
	Flag               string     `json:"flag"`                  //属性，0 免费 1 热门 2 会员 3 推荐
	Chapter_price      int        `json:"chapter_price"`         //章节定价
	Free_chapter       int        `json:"free_chapter"`          //免费章节
	Chapter_updated_at *time.Time `json:"chapter_updated_at"`    //最新章节时间
	Views              int        `json:"views"`                 //浏览次数
	Collect_num        int        `json:"collect_num"`           //收藏数
	Online_status      int        `json:"online_status"`         //上架状态：0 已上架 1 已下架
	Is_sensitivity     int        `json:"is_sensitivity"`        //是否敏感： 0 不敏感 1 敏感
	CreatedAt          time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt          *time.Time `json:"deleted_at"`
}

func (C *Comics) TableName() string {
	return "comics"
}

func init() {
	//db.Register(&Comics{})
}
