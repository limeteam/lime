package dao

import (
	"fmt"
	"lime/pkg/api/front/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type BooksDao struct {}

func (c BooksDao) Get(id int) model.Books {
	var Books model.Books
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Books)
	return Books
}

func (c BooksDao) List(listDto dto.GeneralListDto) ([]model.Books, int64) {
	var Books []model.Books
	var total int64
	db := db.GetGormDB()
	for sk, sv := range dto.TransformSearch(listDto.Q, dto.BookListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Books)
	db.Model(&model.Books{}).Count(&total)
	return Books, total
}

func (c BooksDao) GetRandBooks(extraId int) []model.Books {
	var Books []model.Books
	db := db.GetGormDB()
	db.Model(&model.Books{}).Where("id !=?",extraId).Order("created_at desc").Limit(2).Find(&Books)
	return Books
}