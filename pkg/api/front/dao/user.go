package dao

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

func (c UsersDao) GetUserByEmailAndPassword(username string, password string) (*model.Users, map[string]string) {
	var user model.Users
	db := db.GetGormDB()
	dbErr := map[string]string{}
	err := db.Debug().Where("username = ?", username).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password
	err = VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}

func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}