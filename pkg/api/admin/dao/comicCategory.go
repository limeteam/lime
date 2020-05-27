package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type ComicCategoryDao struct {
}

func (c ComicCategoryDao) Get(id int) model.ComicCategory {
	var Category model.ComicCategory
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Category)
	return Category
}

func (c ComicCategoryDao) GetAll() []model.ComicCategory {
	var Category []model.ComicCategory
	db := db.GetGormDB()
	db.Model(&model.ComicCategory{}).Find(&Category)
	return Category
}

func (c ComicCategoryDao) List(listDto dto.ComicCategoryListDto) ([]model.ComicCategory, int64) {
	var Category []model.ComicCategory
	var total int64
	db := db.GetGormDB()
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Category)
	db.Model(&model.ComicCategory{}).Count(&total)
	return Category, total
}

func (c ComicCategoryDao) Create(Category *model.ComicCategory) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Category)
}

// Update
func (c ComicCategoryDao) Update(Category *model.ComicCategory, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Category).Update(ups)
}

func (c ComicCategoryDao) Delete(Category *model.ComicCategory) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Category)
}
