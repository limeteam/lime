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
	//ChannelId, err := strconv.Atoi(listDto.ChannelId)
	db := db.GetGormDB()
	//if err == nil && ChannelId == 1 {
	//	db = db.Where("created_at >=?", listDto.Start_time)
	//	db = db.Where("created_at <=?", listDto.End_time)
	//}
	//db.Where("deleted_at =?", 0)
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Category)
	db.Model(&model.Category{}).Count(&total)
	return Category, total
}

func (c CategoryDao) Create(Category *model.Category) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Category)
}

func (c CategoryDao) Update(Category *model.Category) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Category)
}

func (c CategoryDao) Delete(Category *model.Category) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Category)
}
