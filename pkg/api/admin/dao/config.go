package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type ConfigDao struct {
}

func (c ConfigDao) Get(id int) model.Config {
	var Config model.Config
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Config)
	return Config
}

func (c ConfigDao) GetByCode(code string) model.Config {
	var Config model.Config
	db := db.GetGormDB()
	db.Where("config_code = ?", code).First(&Config)
	return Config
}

func (c ConfigDao) GetAll() []model.Config {
	var Config []model.Config
	db := db.GetGormDB()
	db.Model(&model.Config{}).Find(&Config)
	return Config
}

func (c ConfigDao) List(listDto dto.GeneralListDto) ([]model.Config, int64) {
	var Config []model.Config
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
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Config)
	db.Model(&model.Config{}).Count(&total)
	return Config, total
}

func (c ConfigDao) Create(Config *model.Config) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Config)
}

// Update
func (c ConfigDao) Update(Config *model.Config, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Config).Update(ups)
}

func (c ConfigDao) Delete(Config *model.Config) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Config)
}
