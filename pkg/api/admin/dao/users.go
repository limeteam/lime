package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type UsersDao struct {
}

func (c UsersDao) Get(id int) model.Users {
	var Users model.Users
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Users)
	return Users
}

func (c UsersDao) GetAll() []model.Users {
	var Users []model.Users
	db := db.GetGormDB()
	db.Model(&model.Users{}).Find(&Users)
	return Users
}

func (c UsersDao) List(listDto dto.GeneralListDto) ([]model.Users, int64) {
	var Users []model.Users
	var total int64
	db := db.GetGormDB()
	for sk, sv := range dto.TransformSearch(listDto.Q, dto.BookListSearchMapping) {
		if sk == "name" {
			db = db.Where("username like ?", "%"+sv+"%").Or("mobile = ?", sv).Or("id = ?", sv)
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
		}
	}
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Users)
	db.Model(&model.Users{}).Count(&total)
	return Users, total
}

func (c UsersDao) Create(Users *model.Users) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Users)
}

// Update
func (c UsersDao) Update(Users *model.Users, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(Users).Update(ups)
}

func (c UsersDao) Delete(Users *model.Users) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Users)
}
