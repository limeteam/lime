package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"lime/pkg/common/db"
)

type ComicCommentDao struct {
}

func (c ComicCommentDao) Get(id int) model.ComicComments {
	var Comment model.ComicComments
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Comment)
	return Comment
}

func (c ComicCommentDao) GetAll() []model.ComicComments {
	var Comment []model.ComicComments
	db := db.GetGormDB()
	db.Model(&model.ComicComments{}).Find(&Comment)
	return Comment
}

func (c ComicCommentDao) List(listDto dto.ComicCommentListDto) ([]model.ComicComments, int64) {
	var Comment []model.ComicComments
	var total int64
	db := db.GetGormDB()
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Comment)
	db.Model(&model.ComicComments{}).Count(&total)
	return Comment, total
}

func (c ComicCommentDao) Create(Comment *model.ComicComments) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Comment)
}

func (c ComicCommentDao) Delete(Comment *model.ComicComments) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Comment)
}
