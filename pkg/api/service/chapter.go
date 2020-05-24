package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/dao"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
)

var chapterDao = dao.ChapterDao{}

// ChapterService
type ChapterService struct {
}

// InfoOfId - get info by id
func (us ChapterService) InfoOfId(dto dto.GeneralGetDto) model.Chapter {
	return chapterDao.Get(dto.Id)
}

// List - users list with pagination
func (us ChapterService) List(dto dto.GeneralListDto) ([]model.Chapter, int64) {
	return chapterDao.List(dto)
}

// Create - create a new Chapter
func (us ChapterService) Create(dto dto.ChapterCreateDto) model.Chapter {
	ChapterModel := model.Chapter{
		Book_id: dto.Book_id,
		Title:   dto.Title,
		Content: dto.Content,
		Status:  dto.Status,
		Url:     dto.Url,
	}
	c := chapterDao.Create(&ChapterModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return ChapterModel
}

// Update - update Chapter's information
func (us ChapterService) Update(dto dto.ChapterEditDto) int64 {
	ChapterModel := model.Chapter{
		Id:      dto.Id,
		Book_id: dto.Book_id,
		Title:   dto.Title,
		Content: dto.Content,
		Status:  dto.Status,
		Url:     dto.Url,
	}

	c := chapterDao.Update(&ChapterModel)
	return c.RowsAffected
}

// Delete - delete Chapter
func (us ChapterService) Delete(dto dto.GeneralDelDto) int64 {
	ChapterModel := model.Chapter{
		Id: dto.Id,
	}
	c := chapterDao.Delete(&ChapterModel)
	return c.RowsAffected
}

