package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/admin/dao"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"time"
)

var DistributorDao = dao.DistributorDao{}

// Service
type DistributorService struct {
}

// InfoOfId
func (cs DistributorService) InfoOfId(dto dto.GeneralGetDto) model.Distributor {
	return DistributorDao.Get(dto.Id)
}

func (cs DistributorService) GetAll() []model.Distributor {
	return DistributorDao.GetAll()
}

// List
func (cs DistributorService) List(dto dto.GeneralListDto) ([]model.Distributor, int64) {
	return DistributorDao.List(dto)
}

// Create
func (cs DistributorService) Create(dto dto.DistributorCreateDto) (model.Distributor, error) {
	Model := model.Distributor{
		Recommender_id:       dto.Recommender_id,
		Distributor_id:       dto.Distributor_id,
		Name:                 dto.Name,
		Mobile:               dto.Mobile,
		Distribution_level:   dto.Distribution_level,
		Commission_withdrawn: dto.Commission_withdrawn,
		Commission_available: dto.Commission_available,
		Status:               dto.Status,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	c := DistributorDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (cs DistributorService) Update(dto dto.DistributorEditDto) int64 {
	Model := model.Distributor{
		Id: dto.Id,
	}
	c := DistributorDao.Update(&Model, map[string]interface{}{
		"recommender_id":       dto.Recommender_id,
		"distributor_id":       dto.Distributor_id,
		"name":                 dto.Name,
		"mobile":               dto.Mobile,
		"distribution_level":   dto.Distribution_level,
		"commission_withdrawn": dto.Commission_withdrawn,
		"commission_available": dto.Commission_available,
		"status":               dto.Status,
		"updated_at":           time.Now(),
	})
	return c.RowsAffected
}

// Update
func (cs DistributorService) UpdateStatus(dto dto.DistributorUpdateStatusDto) int64 {
	Model := model.Distributor{
		Id: dto.Id,
	}
	c := DistributorDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (cs DistributorService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.Distributor{
		Id: dto.Id,
	}
	c := DistributorDao.Delete(&Model)
	return c.RowsAffected
}
