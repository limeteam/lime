package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type CommentDao struct {
}

func (c CommentDao) Get(id int) model.Comments {
	var Comment model.Comments
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Comment)
	return Comment
}

func (c CommentDao) GetAll() []model.Comments {
	var Comment []model.Comments
	db := db.GetGormDB()
	db.Model(&model.Comments{}).Find(&Comment)
	return Comment
}

func (c CommentDao) List(listDto dto.CommentListDto) ([]model.Comments, int64) {
	var Comment []model.Comments
	var total int64
	db := db.GetGormDB()
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Comment)
	db.Model(&model.Comments{}).Count(&total)
	return Comment, total
}

func (c CommentDao) Create(Comment *model.Comments) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Comment)
}

func (c CommentDao) Delete(Comment *model.Comments) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Comment)
}
