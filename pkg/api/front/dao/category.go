package dao

import (
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type CategoryDao struct {}


func (c CategoryDao) GetAll() []model.Category {
	var Category []model.Category
	db := db.GetGormDB()
	db.Model(&model.Category{}).Find(&Category)
	return Category
}
