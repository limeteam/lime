package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/dao"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"time"
)

var ComicCategoryDao = dao.ComicCategoryDao{}

// Service
type ComicCategoryService struct {
}

// InfoOfId
func (cs ComicCategoryService) InfoOfId(dto dto.GeneralGetDto) model.ComicCategory {
	return ComicCategoryDao.Get(dto.Id)
}

func (cs ComicCategoryService) GetAll() []model.ComicCategory {
	return ComicCategoryDao.GetAll()
}

// List
func (cs ComicCategoryService) List(dto dto.ComicCategoryListDto) ([]model.ComicCategory, int64) {
	return ComicCategoryDao.List(dto)
}

// Create
func (cs ComicCategoryService) Create(dto dto.ComicCategoryCreateDto) (model.ComicCategory, error) {
	ComicCategoryModel := model.ComicCategory{
		Name:      dto.Name,
		Comic_num:  dto.Comic_num,
		Sort:      dto.Sort,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	c := ComicCategoryDao.Create(&ComicCategoryModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return ComicCategoryModel, nil
}

// Update
func (cs ComicCategoryService) Update(dto dto.ComicCategoryEditDto) int64 {
	ComicCategoryModel := model.ComicCategory{
		Id:        dto.Id,
		Name:      dto.Name,
		Comic_num: dto.Comic_num,
		Sort:      dto.Sort,
	}
	c := ComicCategoryDao.Update(&ComicCategoryModel, map[string]interface{}{
		"name":       dto.Name,
		"comic_num": dto.Comic_num,
		"sort":       dto.Sort,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (cs ComicCategoryService) Delete(dto dto.GeneralDelDto) int64 {
	ComicCategoryModel := model.ComicCategory{
		Id: dto.Id,
	}
	c := ComicCategoryDao.Delete(&ComicCategoryModel)
	return c.RowsAffected
}
