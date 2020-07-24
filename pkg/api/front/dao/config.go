package dao

import (
	"gorm.io/datatypes"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/cache"
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

func (c ConfigDao) GetByCode(code string) datatypes.JSON {
	var Config model.Config
	configCache, err := cache.Get(code)
	if err == nil {
		return datatypes.JSON(configCache)
	}
	db := db.GetGormDB()
	db.Where("config_code = ?", code).First(&Config)
	cache.Set(code, string(Config.Config_value), 3600)
	return Config.Config_value
}
