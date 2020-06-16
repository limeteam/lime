package service

import (
	"github.com/spf13/viper"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	)

// Service
type UploadService struct {
}

func (us UploadService) GetToken() (token string) {
	accessKey := viper.GetString("qiniu.ak")
	secretKey := viper.GetString("qiniu.sk")
	bucket:= viper.GetString("qiniu.bucket")
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	return putPolicy.UploadToken(mac)
}
