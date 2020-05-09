package dto

import "time"

type CategoryListDto struct {
	Name      string    `form:"name" json:"name" `
	ChannelId string    `form:"channel_id" json:"channel_id" `
	NovelNum  string    `form:"novel_num,default=0" json:"novel_num" `
	Sort      int       `form:"sort" json:"sort" `
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
	DeletedAt time.Time `form:"deleted_at" json:"deleted_at"`
	Page      string    `form:"page" json:"page" `
	Skip      int       `form:"skip,default=0" json:"skip"`
	Limit     int       `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type CategoryCreateDto struct {
	Name      string    `form:"name" json:"name" `
	ChannelId string    `form:"channel_id" json:"channel_id" `
	NovelNum  string    `form:"novel_num,default=0" json:"novel_num" `
	Sort      int       `form:"sort" json:"sort" `
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
	DeletedAt time.Time `form:"deleted_at" json:"deleted_at"`
}

type CategoryEditDto struct {
	Name      string    `form:"name" json:"name" `
	ChannelId string    `form:"channel_id" json:"channel_id" `
	NovelNum  string    `form:"novel_num,default=0" json:"novel_num" `
	Sort      int       `form:"sort" json:"sort" `
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
	DeletedAt time.Time `form:"deleted_at" json:"deleted_at"`
}
