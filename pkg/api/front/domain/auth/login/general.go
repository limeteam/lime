package login

import (
	"lime/pkg/api/front/domain/auth"
	"lime/pkg/api/admin/model"
)

// VerifyPassword : verify password by salt
func VerifyPassword(password string, userModel model.Users) bool {
	if pwd, err := auth.HashPassword(password, userModel.Salt); err == nil && pwd == userModel.Password {
		return true
	}
	return false
}
