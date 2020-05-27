package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type ComicsDao struct {
}

func (c ComicsDao) Get(id int) model.Comics {
	var Comics model.Comics
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Comics)
	return Comics
}

func (c ComicsDao) GetAll() []model.Comics {
	var Comics []model.Comics
	db := db.GetGormDB()
	db.Model(&model.Comics{}).Find(&Comics)
	return Comics
}

func (c ComicsDao) List(listDto dto.GeneralListDto) ([]model.Comics, int64) {
	var Comics []model.Comics
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
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Comics)
	db.Model(&model.Comics{}).Count(&total)
	return Comics, total
}

func (c ComicsDao) Create(Comics *model.Comics) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Comics)
}

// Update
func (c ComicsDao) Update(Comics *model.Comics, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Comics).Update(ups)
}

func (c ComicsDao) Delete(Comics *model.Comics) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Comics)
}
