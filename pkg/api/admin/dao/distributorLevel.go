package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type DistributorLevelDao struct {
}

func (c DistributorLevelDao) Get(id int) model.DistributorLevel {
	var DistributorLevel model.DistributorLevel
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&DistributorLevel)
	return DistributorLevel
}

func (c DistributorLevelDao) GetAll() []model.DistributorLevel {
	var DistributorLevel []model.DistributorLevel
	db := db.GetGormDB()
	db.Model(&model.DistributorLevel{}).Find(&DistributorLevel)
	return DistributorLevel
}

func (c DistributorLevelDao) List(listDto dto.GeneralListDto) ([]model.DistributorLevel, int64) {
	var DistributorLevel []model.DistributorLevel
	var total int64
	db := db.GetGormDB()
	for sk, sv := range dto.TransformSearch(listDto.Q, dto.BookListSearchMapping) {
		if sk == "name" {
			db = db.Where("name like ?", "%"+sv+"%").Or("author = ?", sv).Or("source = ?", sv)
		} else if sk == "flag" {
			db = db.Where("FIND_IN_SET(?, "+sk+")", sv)
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
		}
	}
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&DistributorLevel)
	db.Model(&model.DistributorLevel{}).Count(&total)
	return DistributorLevel, total
}

func (c DistributorLevelDao) Create(DistributorLevel *model.DistributorLevel) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(DistributorLevel)
}

// Update
func (c DistributorLevelDao) Update(DistributorLevel *model.DistributorLevel, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(DistributorLevel).Update(ups)
}

func (c DistributorLevelDao) Delete(DistributorLevel *model.DistributorLevel) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(DistributorLevel)
}
