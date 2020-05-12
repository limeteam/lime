package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
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

func (c BooksDao) List(listDto dto.BooksListDto) ([]model.Books, int64) {
	var Books []model.Books
	var total int64
	//ChannelId, err := strconv.Atoi(listDto.ChannelId)
	db := db.GetGormDB()
	//if err == nil && ChannelId == 1 {
	//	db = db.Where("created_at >=?", listDto.Start_time)
	//	db = db.Where("created_at <=?", listDto.End_time)
	//}
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
