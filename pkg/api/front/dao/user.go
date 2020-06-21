package dao

import (
	"github.com/jinzhu/gorm"
	"lime/pkg/api/admin/model"
	"lime/pkg/api/front/domain/auth/login"
	"lime/pkg/api/utils"
	"lime/pkg/common/db"
)

type UsersDao struct{}

func (c UsersDao) GetUserByField(username string,field string) model.Users {
	var Users model.Users
	db := db.GetGormDB()
	db.Where( field +"= ?", username).First(&Users)
	return Users
}

func (c UsersDao) GetUserByUsernameAndPassword(username string, password string) (*model.Users, map[string]string) {
	var user model.Users
	db := db.GetGormDB()
	dbErr := map[string]string{}
	var err error
	if utils.VerifyEmailFormat(username){
		err = db.Debug().Where("email = ?", username).Take(&user).Error
	}else if utils.VerifyMobileFormat(username) {
		err = db.Debug().Where("mobile = ?", username).Take(&user).Error
	}else {
		err = db.Debug().Where("username = ?", username).Take(&user).Error
	}

	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password
	if !login.VerifyPassword(password, user) {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}

//Get - get single user info
func (u UsersDao) Get(id int) model.Users {
	var user model.Users
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&user)
	return user
}

func (u UsersDao) Create(Users *model.Users) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Users)
}

// Update - update user
func (u UsersDao) Update(user *model.Users, ups map[string]interface{}) *gorm.DB {
	db := db.GetGormDB()
	return db.Model(user).Update(ups)
}