package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"lime/pkg/common/db"
)

type CategoryDao struct {
}

func (c CategoryDao) Get(id int) model.Category {
	var Category model.Category
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Category)
	return Category
}

func (c CategoryDao) GetAll() []model.Category {
	var Category []model.Category
	db := db.GetGormDB()
	db.Model(&model.Category{}).Find(&Category)
	return Category
}

func (c CategoryDao) List(listDto dto.CategoryListDto) ([]model.Category, int64) {
	var Category []model.Category
	var total int64
	db := db.GetGormDB()
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Category)
	db.Model(&model.Category{}).Count(&total)
	return Category, total
}

func (c CategoryDao) Create(Category *model.Category) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Category)
}

// Update
func (c CategoryDao) Update(Category *model.Category, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Category).Update(ups)
}

func (c CategoryDao) Delete(Category *model.Category) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Category)
}
