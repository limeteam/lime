package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type DistributorDao struct {
}

func (c DistributorDao) Get(id int) model.Distributor {
	var Distributor model.Distributor
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Distributor)
	return Distributor
}

func (c DistributorDao) GetAll() []model.Distributor {
	var Distributor []model.Distributor
	db := db.GetGormDB()
	db.Model(&model.Distributor{}).Find(&Distributor)
	return Distributor
}

func (c DistributorDao) List(listDto dto.GeneralListDto) ([]model.Distributor, int64) {
	var Distributor []model.Distributor
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
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Distributor)
	db.Model(&model.Distributor{}).Count(&total)
	return Distributor, total
}

func (c DistributorDao) Create(Distributor *model.Distributor) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Distributor)
}

// Update
func (c DistributorDao) Update(Distributor *model.Distributor, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Distributor).Update(ups)
}

func (c DistributorDao) Delete(Distributor *model.Distributor) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Distributor)
}
