package dao

import (
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type ChaptersDao struct {
}

func (c ChaptersDao) Get(id int) model.Chapters {
	var Chapters model.Chapters
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Chapters)
	return Chapters
}

func (c ChaptersDao) GetAll() []model.Chapters {
	var Chapters []model.Chapters
	db := db.GetGormDB()
	db.Model(&model.Chapters{}).Find(&Chapters)
	return Chapters
}
