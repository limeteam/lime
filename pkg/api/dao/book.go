package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"lime/pkg/common/db"
)

type BookDao struct {
}

func (u BookDao) Get(id int) model.Book {
	var Book model.Book
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Book)
	return Book
}

func (u BookDao) List(listDto dto.BookListDto) ([]model.Book, int64) {
	var Book []model.Book
	var total int64
	db := db.GetGormDB()
	db.Preload("Book").Offset(listDto.Skip).Limit(listDto.Limit).Find(&Book)
	db.Model(&model.Book{}).Count(&total)
	return Book, total
}

func (u BookDao) Create(Book *model.Book) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Book)
}

func (u BookDao) Update(Book *model.Book) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Book)
}

func (u BookDao) Delete(Book *model.Book) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Book)
}
