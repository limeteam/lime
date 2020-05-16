package model

import (
	"lime/pkg/common/db"
	"time"
)

type Books struct {
	Id                        int        `gorm:"primary_key" json:"id"`     //分类ID
	Name                      string     `json:"name"`                      //分类名称
	Old_name                  string     `json:"old_name"`                  //原名
	Channel_id                int        `json:"channel_id"`                //所属频道
	Category_id               int        `json:"category_id"`               //小说分类
	Desc                      string     `json:"desc"`                      //小说描述
	Cover                     string     `json:"cover"`                     //小说封面
	Author                    string     `json:"author"`                    //小说作者
	Source                    string     `json:"source"`                    //来源
	Split_rule                int        `json:"split_rule"`                //拆分规则：0 智能分章 1 固定字数分章 2 标签分章
	Upload_file               string     `json:"upload_file"`               //上传小说
	Status                    int        `json:"status"`                    //小说状态，0连载中，1已完结，2太监
	Flag                      string     `json:"flag"`                      //属性，0 新书 1 热门 2 会员 3 热门 4 精选 5 大神
	Chapter_price             int        `json:"chapter_price"`             //每章价格
	Thousand_characters_price int        `json:"thousand_characters_price"` //千字价格
	Score                     int        `json:"score"`                     //评分
	Text_num                  int        `json:"text_num"`                  //小说字数
	Chapter_num               int        `json:"chapter_num"`               //小说章节数
	Chapter_updated_at        *time.Time `json:"chapter_updated_at"`        //最新章节时间
	Chapter_id                int        `json:"chapter_id"`                //最新章节id
	Chapter_title             string     `json:"chapter_title"`             //最新章节标题
	Views                     int        `json:"views"`                     //浏览次数
	Collect_num               int        `json:"collect_num"`               //收藏数
	Online_status             int        `json:"online_status"`             //上架状态：0 已上架 1 已下架
	Is_sensitivity            int        `json:"is_sensitivity"`            //是否敏感： 0 不敏感 1 敏感
	CreatedAt                 time.Time  `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt                 time.Time  `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt                 *time.Time `json:"deleted_at"`
}

func (B *Books) TableName() string {
	return "novel_book"
}

func init() {
	db.Register(&Books{})
}
