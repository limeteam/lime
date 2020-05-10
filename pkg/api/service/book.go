package service

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/dao"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
)

var BookDao = dao.BookDao{}

// Service
type BookService struct{}

// InfoOfId
func (us BookService) InfoOfId(dto dto.GeneralGetDto) model.Book {
	return BookDao.Get(dto.Id)
}

// List
func (us BookService) List(dto dto.BookListDto) ([]model.Book, int64) {
	return BookDao.List(dto)
}

// Create
func (us BookService) Create(dto dto.BookCreateDto) model.Book {
	BookModel := model.Book{
		Name:   dto.Name,
		Author: dto.Author,
		Image:  dto.Image,
		Status: dto.Status,
		Url:    dto.Url,
		Desc:   dto.Desc,
	}
	c := BookDao.Create(&BookModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return BookModel
}

// Update
func (us BookService) Update(dto dto.BookEditDto) int64 {
	BookModel := BookDao.Get(dto.Id)
	BookModel.Name = dto.Name
	BookModel.Author = dto.Author
	BookModel.Image = dto.Image
	BookModel.Status = dto.Status
	BookModel.Url = dto.Url
	BookModel.Desc = dto.Desc
	c := BookDao.Update(&BookModel)
	return c.RowsAffected
}

// Delete
func (us BookService) Delete(dto dto.GeneralDelDto) int64 {
	BookModel := model.Book{
		Id: dto.Id,
	}
	c := BookDao.Delete(&BookModel)
	return c.RowsAffected
}
