package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/admin/dao"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"time"
)

var ComicCommentDao = dao.ComicCommentDao{}

// Service
type ComicCommentervice struct {
}

// InfoOfId
func (cs ComicCommentervice) InfoOfId(dto dto.GeneralGetDto) model.ComicComments {
	return ComicCommentDao.Get(dto.Id)
}

func (cs ComicCommentervice) GetAll() []model.ComicComments {
	return ComicCommentDao.GetAll()
}

// List
func (cs ComicCommentervice) List(dto dto.ComicCommentListDto) ([]model.ComicComments, int64) {
	return ComicCommentDao.List(dto)
}

// Create
func (cs ComicCommentervice) Create(dto dto.ComicCommentCreateDto) (model.ComicComments, error) {
	ComicCommentModel := model.ComicComments{
		Comic_id:  dto.Comic_id,
		Username:  dto.Username,
		Content:   dto.Content,
		Source:    dto.Source,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	c := ComicCommentDao.Create(&ComicCommentModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return ComicCommentModel, nil
}

//Delete
func (cs ComicCommentervice) Delete(dto dto.GeneralDelDto) int64 {
	ComicCommentModel := model.ComicComments{
		Id: dto.Id,
	}
	c := ComicCommentDao.Delete(&ComicCommentModel)
	return c.RowsAffected
}
