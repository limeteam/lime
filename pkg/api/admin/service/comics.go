package service

import (
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"io"
	"lime/pkg/api/admin/dao"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var ComicsDao = dao.ComicsDao{}

// Service
type ComicsService struct {
}

// InfoOfId
func (cs ComicsService) InfoOfId(dto dto.GeneralGetDto) model.Comics {
	return ComicsDao.Get(dto.Id)
}

func (cs ComicsService) GetAll() []model.Comics {
	return ComicsDao.GetAll()
}

// List
func (cs ComicsService) List(dto dto.GeneralListDto) ([]model.Comics, int64) {
	return ComicsDao.List(dto)
}

// Create
func (cs ComicsService) Create(dto dto.ComicsCreateDto) (model.Comics, error) {
	Model := model.Comics{
		Name:               dto.Name,
		Old_name:           dto.Old_name,
		Channel_id:         dto.Channel_id,
		Category_id:        dto.Category_id,
		Desc:               dto.Desc,
		Horizontal_cover:   dto.Horizontal_cover,
		Vertical_cover:     dto.Vertical_cover,
		Author:             dto.Author,
		Source:             dto.Source,
		Status:             dto.Status,
		Flag:               dto.Flag,
		Chapter_price:      dto.Chapter_price,
		Chapter_updated_at: dto.Chapter_updated_at,
		Views:              dto.Views,
		Collect_num:        dto.Collect_num,
		Online_status:      dto.Online_status,
		Is_sensitivity:     dto.Is_sensitivity,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	c := ComicsDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (cs ComicsService) Update(dto dto.ComicsEditDto) int64 {
	Model := model.Comics{
		Id:                 dto.Id,
		Name:               dto.Name,
		Old_name:           dto.Old_name,
		Channel_id:         dto.Channel_id,
		Category_id:        dto.Category_id,
		Desc:               dto.Desc,
		Horizontal_cover:   dto.Horizontal_cover,
		Vertical_cover:     dto.Vertical_cover,
		Author:             dto.Author,
		Source:             dto.Source,
		Status:             dto.Status,
		Flag:               dto.Flag,
		Chapter_price:      dto.Chapter_price,
		Chapter_updated_at: dto.Chapter_updated_at,
		Views:              dto.Views,
		Collect_num:        dto.Collect_num,
		Online_status:      dto.Online_status,
		Is_sensitivity:     dto.Is_sensitivity,
		UpdatedAt:          time.Now(),
	}
	c := ComicsDao.Update(&Model, map[string]interface{}{
		"name":               dto.Name,
		"old_name":           dto.Old_name,
		"channel_id":         dto.Channel_id,
		"category_id":        dto.Category_id,
		"desc":               dto.Desc,
		"horizontal_cover":   dto.Horizontal_cover,
		"vertical_cover":     dto.Vertical_cover,
		"author":             dto.Author,
		"source":             dto.Source,
		"status":             dto.Status,
		"flag":               dto.Flag,
		"chapter_price":      dto.Chapter_price,
		"chapter_updated_at": dto.Chapter_updated_at,
		"views":              dto.Views,
		"collect_num":        dto.Collect_num,
		"online_status":      dto.Online_status,
		"is_sensitivity":     dto.Is_sensitivity,
		"updated_at":         time.Now(),
	})
	return c.RowsAffected
}

// Update
func (cs ComicsService) UpdateStatus(dto dto.ComicsUpdateStatusDto) int64 {
	Model := model.Comics{
		Id: dto.Id,
	}
	c := ComicsDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (cs ComicsService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.Comics{
		Id: dto.Id,
	}
	c := ComicsDao.Delete(&Model)
	return c.RowsAffected
}

//上传封面
func (cs ComicsService) UploadCover(file multipart.File, filename string) (filepath string, err error) {
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

//上传封面
func (cs ComicsService) UploadComicsFile(file multipart.File, filename string) (filepath string, err error) {
	defer file.Close()
	filenameWithSuffix := path.Base(filename)
	fileSuffix := path.Ext(filenameWithSuffix)
	filename = strconv.FormatInt(time.Now().Unix(), 10) + fileSuffix
	destFile, err := os.Create("./data/Comics/" + filename)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return
	}
	filepath = "/upload/Comics/" + filename
	return filepath, nil
}
