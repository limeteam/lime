package service

import (
	"lime/pkg/api/dao"
	//"lime/pkg/api/domain/category"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	//log "github.com/sirupsen/logrus"
)

var CategoryDao = dao.CategoryDao{}

// Service
type CategoryService struct {
}

// InfoOfId
func (c CategoryService) InfoOfId(dto dto.GeneralGetDto) model.Category {
	return CategoryDao.Get(dto.Id)
}

func (c CategoryService) GetAll() []model.Category {
	return CategoryDao.GetAll()
}

// List
func (c CategoryService) List(dto dto.CategoryListDto) ([]model.Category, int64) {
	return CategoryDao.List(dto)
}

// Create
//func (c CategoryService) Create(dto dto.CategoryCreateDto) (model.Category, error) {
//	CategoryModel := model.Category{
//		Name:          dto.Name,
//		ChannelId:  dto.ChannelId,
//		NovelNum: dto.NovelNum,
//		Sort: dto.Sort,
//		CreatedAt: dto.CreatedAt,
//		UpdatedAt: dto.UpdatedAt,
//		DeletedAt: dto.DeletedAt,
//	}
//	c := CategoryDao.Create(&CategoryModel)
//	if c.Error != nil {
//		log.Error(c.Error.Error())
//	}
//
//	return CategoryModel, nil
//}

// Update
func (c CategoryService) Update(dto dto.TaskEditDto) error {
	return nil
}

// Delete
//func (c CategoryService) Delete(dto dto.GeneralDelDto) int64 {
//	CategoryModel := model.Category{
//		Id: dto.Id,
//	}
//	c := CategoryDao.Delete(&CategoryModel)
//	return c.RowsAffected
//}
