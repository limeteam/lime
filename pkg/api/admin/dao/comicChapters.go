package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type ComicChaptersDao struct {
}

func (c ComicChaptersDao) Get(id int) model.ComicChapters {
	var ComicChapters model.ComicChapters
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&ComicChapters)
	return ComicChapters
}

func (c ComicChaptersDao) GetAll() []model.ComicChapters {
	var ComicChapters []model.ComicChapters
	db := db.GetGormDB()
	db.Model(&model.ComicChapters{}).Find(&ComicChapters)
	return ComicChapters
}

func (c ComicChaptersDao) List(listDto dto.GeneralListDto) ([]model.ComicChapters, int64) {
	var ComicChapters []model.ComicChapters
	var total int64
	db := db.GetGormDB()
	for sk, sv := range dto.TransformSearch(listDto.Q, dto.ChaptersListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&ComicChapters)
	db.Model(&model.ComicChapters{}).Count(&total)
	return ComicChapters, total
}

func (c ComicChaptersDao) Create(ComicChapters *model.ComicChapters) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(ComicChapters)
}

// Update
func (c ComicChaptersDao) Update(ComicChapters *model.ComicChapters, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(ComicChapters).Update(ups)
}

func (c ComicChaptersDao) Delete(ComicChapters *model.ComicChapters) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(ComicChapters)
}
