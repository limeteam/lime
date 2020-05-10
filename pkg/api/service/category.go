package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/dao"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"time"
)

var CategoryDao = dao.CategoryDao{}

// Service
type CategoryService struct {
}

// InfoOfId
func (cs CategoryService) InfoOfId(dto dto.GeneralGetDto) model.Category {
	return CategoryDao.Get(dto.Id)
}

func (cs CategoryService) GetAll() []model.Category {
	return CategoryDao.GetAll()
}

// List
func (cs CategoryService) List(dto dto.CategoryListDto) ([]model.Category, int64) {
	return CategoryDao.List(dto)
}

// Create
func (cs CategoryService) Create(dto dto.CategoryCreateDto) (model.Category, error) {
	CategoryModel := model.Category{
		Name:      dto.Name,
		ChannelId: dto.ChannelId,
		NovelNum:  dto.NovelNum,
		Sort:      dto.Sort,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	c := CategoryDao.Create(&CategoryModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return CategoryModel, nil
}

// Update
func (cs CategoryService) Update(dto dto.CategoryEditDto) int64 {
	categoryModel := model.Category{
		Id:        dto.Id,
		Name:      dto.Name,
		ChannelId: dto.ChannelId,
		Sort:      dto.Sort,
		NovelNum:  dto.NovelNum,
	}
	c := CategoryDao.Update(&categoryModel, map[string]interface{}{
		"name":       dto.Name,
		"channel_id": dto.ChannelId,
		"sort":       dto.Sort,
		"novel_num":  dto.NovelNum,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (cs CategoryService) Delete(dto dto.GeneralDelDto) int64 {
	CategoryModel := model.Category{
		Id: dto.Id,
	}
	c := CategoryDao.Delete(&CategoryModel)
	return c.RowsAffected
}
