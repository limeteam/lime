package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type ChapterDao struct {
}

func (u ChapterDao) Get(id int) model.Chapter {
	var Chapter model.Chapter
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Chapter)
	return Chapter
}

func (u ChapterDao) List(listDto dto.GeneralListDto) ([]model.Chapter, int64) {
	var Chapter []model.Chapter
	var total int64
	db := db.GetGormDB()
	db.Preload("Chapter").Offset(listDto.Skip).Limit(listDto.Limit).Find(&Chapter)
	db.Model(&model.Chapter{}).Count(&total)
	return Chapter, total
}

func (u ChapterDao) Create(Chapter *model.Chapter) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Chapter)
}

func (u ChapterDao) Update(Chapter *model.Chapter) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Chapter)
}

func (u ChapterDao) Delete(Chapter *model.Chapter) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Chapter)
}
