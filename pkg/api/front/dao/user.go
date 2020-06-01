package dao

import (
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type UsersDao struct{}

func (c UsersDao) GetUserByUsername(username string) model.Users {
	var Users model.Users
	db := db.GetGormDB()
	db.Where("username = ?", username).First(&Users)
	return Users
}
