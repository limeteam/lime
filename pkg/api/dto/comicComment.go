package dto

import "time"

type ComicCommentListDto struct {
	Id        int        `form:"id" json:"id" `
	Comic_id  int        `form:"comic_id" json:"comic_id" `
	Username  string     `form:"username" json:"username" `
	Content   string     `form:"content,default=0" json:"content" `
	Source    string     `form:"source" json:"source" `
	CreatedAt time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `form:"deleted_at" json:"deleted_at"`
	Page      string     `form:"page" json:"page" `
	Skip      int        `form:"skip,default=0" json:"skip"`
	Limit     int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type ComicCommentCreateDto struct {
	Comic_id  int        `form:"comic_id" json:"comic_id" `
	Username  string     `form:"username" json:"username" `
	Content   string     `form:"content,default=0" json:"content" `
	Source    string     `form:"source" json:"source" `
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
}
