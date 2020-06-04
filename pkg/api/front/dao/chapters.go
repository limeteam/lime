package dao

import (
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type ChaptersDao struct {
}

func (c ChaptersDao) Get(bookId int,chapterId int) model.Chapters {
	var Chapters model.Chapters
	db := db.GetGormDB()
	db.Where("novel_id = ? and chapter_no = ?", bookId,chapterId).First(&Chapters)
	return Chapters
}

func (c ChaptersDao) GetListByBookId(bookId int) []model.Chapters {
	var Chapters []model.Chapters
	db := db.GetGormDB()
	db.Model(&model.Chapters{}).Where("novel_id = ?",bookId).Find(&Chapters)
	return Chapters
}

func (c ChaptersDao) GetAll() []model.Chapters {
	var Chapters []model.Chapters
	db := db.GetGormDB()
	db.Model(&model.Chapters{}).Find(&Chapters)
	return Chapters
}
