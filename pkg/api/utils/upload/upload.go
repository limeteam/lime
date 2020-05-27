package upload

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"lime/pkg/api/utils"
	"lime/pkg/api/utils/file"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type UploadSetting struct {
	PrefixUrl       string
	RuntimeRootPath string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExts  []string
}

var AppSetting = &UploadSetting{}

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	AppSetting.ImageSavePath = viper.GetString("upload.image_save_path")
	return AppSetting.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	AppSetting.RuntimeRootPath = ""
	return AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	AppSetting.ImageAllowExts = viper.GetStringSlice("upload.image_allow_exts")
	for _, allowExt := range AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		log.Warn(err)
		return false
	}

	return size <= AppSetting.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
