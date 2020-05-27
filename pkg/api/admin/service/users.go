package service

import (
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"lime/pkg/api/admin/dao"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

var UsersDao = dao.UsersDao{}

// Service
type UsersService struct {
}

// InfoOfId
func (bs UsersService) InfoOfId(dto dto.GeneralGetDto) model.Users {
	return UsersDao.Get(dto.Id)
}

func (bs UsersService) GetAll() []model.Users {
	return UsersDao.GetAll()
}

// List
func (bs UsersService) List(dto dto.GeneralListDto) ([]model.Users, int64) {
	return UsersDao.List(dto)
}

// Create
func (bs UsersService) Create(dto dto.UsersCreateDto) (model.Users, error) {
	Model := model.Users{
		Username:        dto.Username,
		Mobile:          dto.Mobile,
		Sex:             dto.Sex,
		Password:        dto.Password,
		Salt:            dto.Salt,
		Faceicon:        dto.Faceicon,
		Wechat:          dto.Wechat,
		Email:           dto.Email,
		Amount:          dto.Amount,
		Coin:            dto.Coin,
		Exempt_login:    dto.Exempt_login,
		Source:          dto.Source,
		Is_vip:          dto.Is_vip,
		Channel_id:      dto.Channel_id,
		Status:          dto.Status,
		CreatedAt:       time.Now(),
		Last_login_time: time.Now(),
	}
	c := UsersDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (bs UsersService) Update(dto dto.UsersEditDto) int64 {
	Model := model.Users{
		Id:              dto.Id,
		Username:        dto.Username,
		Mobile:          dto.Mobile,
		Sex:             dto.Sex,
		Password:        dto.Password,
		Salt:            dto.Salt,
		Faceicon:        dto.Faceicon,
		Wechat:          dto.Wechat,
		Email:           dto.Email,
		Amount:          dto.Amount,
		Coin:            dto.Coin,
		Exempt_login:    dto.Exempt_login,
		Source:          dto.Source,
		Is_vip:          dto.Is_vip,
		Channel_id:      dto.Channel_id,
		Status:          dto.Status,
		Last_login_time: time.Now(),
	}
	c := UsersDao.Update(&Model, map[string]interface{}{
		"Username":        dto.Username,
		"Mobile":          dto.Mobile,
		"Sex":             dto.Sex,
		"Password":        dto.Password,
		"Salt":            dto.Salt,
		"Faceicon":        dto.Faceicon,
		"Wechat":          dto.Wechat,
		"Email":           dto.Email,
		"Amount":          dto.Amount,
		"Coin":            dto.Coin,
		"Exempt_login":    dto.Exempt_login,
		"Source":          dto.Source,
		"Is_vip":          dto.Is_vip,
		"Channel_id":      dto.Channel_id,
		"Status":          dto.Status,
		"Last_login_time": time.Now(),
	})
	return c.RowsAffected
}

// Update
func (bs UsersService) UpdateStatus(dto dto.UsersUpdateStatusDto) int64 {
	Model := model.Users{
		Id: dto.Id,
	}
	c := UsersDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (bs UsersService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.Users{
		Id: dto.Id,
	}
	c := UsersDao.Delete(&Model)
	return c.RowsAffected
}

//上传头像
func (bs UsersService) UploadFace(file multipart.File, filename string) (filepath string, err error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return "", err
	}
	file.Close()
	// 大图
	m := resize.Resize(225, 300, img, resize.Lanczos3)
	filename = strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	out, err := os.Create("./data/images/" + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)

	//缩略图
	dst := "./data/images/" + strings.Replace(filename, ".jpg", "_small.jpg", 1)
	mSmall := resize.Resize(75, 100, img, resize.Lanczos3)
	outSmall, errSmall := os.Create(dst)
	if errSmall != nil {
		return "", errSmall
	}
	defer outSmall.Close()
	jpeg.Encode(outSmall, mSmall, nil)

	filepath = "/upload/images/" + filename
	return filepath, nil
}
