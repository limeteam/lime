package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/admin/dao"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"time"
)

var ConfigDao = dao.ConfigDao{}

// Service
type ConfigService struct {
}

// InfoOfId
func (cs ConfigService) InfoOfId(dto dto.GeneralGetDto) model.Config {
	return ConfigDao.Get(dto.Id)
}

func (cs ConfigService) InfoOfCode(dto dto.ConfigGetByCodeDto) model.Config {
	return ConfigDao.GetByCode(dto.Code)
}

func (cs ConfigService) GetAll() []model.Config {
	return ConfigDao.GetAll()
}

// List
func (cs ConfigService) List(dto dto.GeneralListDto) ([]model.Config, int64) {
	return ConfigDao.List(dto)
}

// Create
func (cs ConfigService) Create(dto dto.ConfigCreateDto) (model.Config, error) {
	Model := model.Config{
		Name:              dto.Name,
		Config_code:       dto.Config_code,
		Config_value:      dto.Config_value,
		Config_group:      dto.Config_group,
		Desc:              dto.Desc,
		Status:            dto.Status,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	c := ConfigDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (cs ConfigService) Update(dto dto.ConfigEditDto) int64 {
	Model := model.Config{
		Id: dto.Id,
	}
	c := ConfigDao.Update(&Model, map[string]interface{}{
		"name":              dto.Name,
		"config_code":       dto.Config_code,
		"config_value":      dto.Config_value,
		"config_group":      dto.Config_group,
		"desc":              dto.Desc,
		"status":            dto.Status,
		"updated_at":        time.Now(),
	})
	return c.RowsAffected
}

// Update
func (cs ConfigService) UpdateStatus(dto dto.ConfigUpdateStatusDto) int64 {
	Model := model.Config{
		Id: dto.Id,
	}
	c := ConfigDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (cs ConfigService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.Config{
		Id: dto.Id,
	}
	c := ConfigDao.Delete(&Model)
	return c.RowsAffected
}

// UpdateByCode
func (cs ConfigService) UpdateByCode(dto dto.ConfigUpdateDto) int64 {
	fmt.Println(dto.Config_code)
	Model := ConfigDao.GetByCode(dto.Config_code)
	if Model.Id <= 0 {
		Model := model.Config{
			Name:              dto.Config_code,
			Config_code:       dto.Config_code,
			Config_value:      dto.Config_value,
			Config_group:      "base_config",
			Desc:              dto.Config_code,
			Status:            1,
			UpdatedAt:         time.Now(),
		}
		c := ConfigDao.Create(&Model)
		if c.Error != nil {
			log.Error(c.Error.Error())
		}
		return c.RowsAffected
	}
	c := ConfigDao.Update(&Model, map[string]interface{}{
		"config_code":       dto.Config_code,
		"config_value":      dto.Config_value,
		"updated_at":        time.Now(),
	})
	return c.RowsAffected
}
