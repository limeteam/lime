package dao

import (
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
	//configCache, err := cache.Get(code)
	//if err == nil {
	//	return configCache
	//}
	db := db.GetGormDB()
	db.Where("config_code = ?", code).First(&Config)
	return Config
}
