package dto

import "time"

type CommentListDto struct {
	Id        int        `form:"id" json:"id" `
	Name      string     `form:"name" json:"name" `
	Novel_id  int        `form:"novel_id" json:"novel_id" `
	Username  string     `form:"username" json:"username" `
	Content   string     `form:"content,default=0" json:"content" `
	Likes     int        `form:"likes" json:"likes" `
	Source    string     `form:"source" json:"source" `
	CreatedAt time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `form:"deleted_at" json:"deleted_at"`
	Page      string     `form:"page" json:"page" `
	Skip      int        `form:"skip,default=0" json:"skip"`
	Limit     int        `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type CommentCreateDto struct {
	Novel_id  int       `form:"novel_id" json:"novel_id" `
	Username  string    `form:"username" json:"username" `
	Content   string    `form:"content,default=0" json:"content" `
	Likes     int       `form:"likes" json:"likes" `
	Source    string    `form:"source" json:"source" `
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
}
