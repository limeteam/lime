package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type BooksDao struct {
}

func (c BooksDao) Get(id int) model.Books {
	var Books model.Books
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Books)
	return Books
}

func (c BooksDao) GetAll() []model.Books {
	var Books []model.Books
	db := db.GetGormDB()
	db.Model(&model.Books{}).Find(&Books)
	return Books
}

func (c BooksDao) List(listDto dto.GeneralListDto) ([]model.Books, int64) {
	var Books []model.Books
	var total int64
	db := db.GetGormDB()
	for sk, sv := range dto.TransformSearch(listDto.Q, dto.BookListSearchMapping) {
		if sk == "name" {
			db = db.Where("name like ?", "%"+sv+"%").Or("author = ?", sv).Or("source = ?", sv)
		} else if sk == "flag" {
			db = db.Where("FIND_IN_SET(?, "+sk+")", sv)
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
		}
	}
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Books)
	db.Model(&model.Books{}).Count(&total)
	return Books, total
}

func (c BooksDao) Create(Books *model.Books) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Books)
}

// Update
func (c BooksDao) Update(Books *model.Books, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Books).Update(ups)
}

func (c BooksDao) Delete(Books *model.Books) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Books)
}
