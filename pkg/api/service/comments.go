package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/dao"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"time"
)

var CommentDao = dao.CommentDao{}

// Service
type CommentService struct {
}

// InfoOfId
func (cs CommentService) InfoOfId(dto dto.GeneralGetDto) model.Comments {
	return CommentDao.Get(dto.Id)
}

func (cs CommentService) GetAll() []model.Comments {
	return CommentDao.GetAll()
}

// List
func (cs CommentService) List(dto dto.CommentListDto) ([]model.Comments, int64) {
	return CommentDao.List(dto)
}

// Create
func (cs CommentService) Create(dto dto.CommentCreateDto) (model.Comments, error) {
	CommentModel := model.Comments{
		Novel_id:  dto.Novel_id,
		Username:  dto.Username,
		Content:   dto.Content,
		Likes:     dto.Likes,
		Source:    dto.Source,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	c := CommentDao.Create(&CommentModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return CommentModel, nil
}

//Delete
func (cs CommentService) Delete(dto dto.GeneralDelDto) int64 {
	CommentModel := model.Comments{
		Id: dto.Id,
	}
	c := CommentDao.Delete(&CommentModel)
	return c.RowsAffected
}
