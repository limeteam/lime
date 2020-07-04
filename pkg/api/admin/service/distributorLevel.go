package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/admin/dao"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"time"
)

var DistributorLevelDao = dao.DistributorLevelDao{}

// Service
type DistributorLevelService struct {
}

// InfoOfId
func (cs DistributorLevelService) InfoOfId(dto dto.GeneralGetDto) model.DistributorLevel {
	return DistributorLevelDao.Get(dto.Id)
}

func (cs DistributorLevelService) GetAll() []model.DistributorLevel {
	return DistributorLevelDao.GetAll()
}

// List
func (cs DistributorLevelService) List(dto dto.GeneralListDto) ([]model.DistributorLevel, int64) {
	return DistributorLevelDao.List(dto)
}

// Create
func (cs DistributorLevelService) Create(dto dto.DistributorLevelCreateDto) (model.DistributorLevel, error) {
	Model := model.DistributorLevel{
		Name:                   dto.Name,
		Repurchase_commission:  dto.Repurchase_commission,
		Rakeback_type:          dto.Rakeback_type,
		Auto_upgrade:           dto.Auto_upgrade,
		Upgrade_conditions:     dto.Upgrade_conditions,
		Adaptive_degradation:   dto.Adaptive_degradation,
		Degradation_conditions: dto.Degradation_conditions,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
	c := DistributorLevelDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (cs DistributorLevelService) Update(dto dto.DistributorLevelEditDto) int64 {
	Model := model.DistributorLevel{
		Id: dto.Id,
	}
	c := DistributorLevelDao.Update(&Model, map[string]interface{}{
		"name":                   dto.Name,
		"repurchase_commission":  dto.Repurchase_commission,
		"recommendtype":          dto.recommendtype,
		"auto_upgrade":           dto.Auto_upgrade,
		"upgrade_conditions":     dto.Upgrade_conditions,
		"adaptive_degradation":   dto.Adaptive_degradation,
		"degradation_conditions": dto.Degradation_conditions,
		"weight":                 dto.Weight,
		"updated_at":             time.Now(),
	})
	return c.RowsAffected
}

// Update
func (cs DistributorLevelService) UpdateStatus(dto dto.DistributorLevelUpdateStatusDto) int64 {
	Model := model.DistributorLevel{
		Id: dto.Id,
	}
	c := DistributorLevelDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (cs DistributorLevelService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.DistributorLevel{
		Id: dto.Id,
	}
	c := DistributorLevelDao.Delete(&Model)
	return c.RowsAffected
}
