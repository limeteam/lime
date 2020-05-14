package service

import (
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"lime/pkg/api/dao"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

var ChaptersDao = dao.ChaptersDao{}

// Service
type ChaptersService struct {
}

// InfoOfId
func (bs ChaptersService) InfoOfId(dto dto.GeneralGetDto) model.Chapters {
	return ChaptersDao.Get(dto.Id)
}

func (bs ChaptersService) GetAll() []model.Chapters {
	return ChaptersDao.GetAll()
}

// List
func (bs ChaptersService) List(dto dto.ChaptersListDto) ([]model.Chapters, int64) {
	return ChaptersDao.List(dto)
}

// Create
func (bs ChaptersService) Create(dto dto.ChaptersCreateDto) (model.Chapters, error) {
	Model := model.Chapters{
		Novel_id:   dto.Novel_id,
		Chapter_no: dto.Chapter_no,
		Title:      dto.Title,
		Desc:       dto.Desc,
		Link:       dto.Link,
		Is_vip:     dto.Is_vip,
		Source:     dto.Source,
		Views:      dto.Views,
		Text_num:   dto.Text_num,
		Status:     dto.Status,
		Try_views:  dto.Try_views,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	c := ChaptersDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (bs ChaptersService) Update(dto dto.ChaptersEditDto) int64 {
	Model := model.Chapters{
		Id:         dto.Id,
		Novel_id:   dto.Novel_id,
		Chapter_no: dto.Chapter_no,
		Title:      dto.Title,
		Desc:       dto.Desc,
		Link:       dto.Link,
		Is_vip:     dto.Is_vip,
		Source:     dto.Source,
		Views:      dto.Views,
		Text_num:   dto.Text_num,
		Status:     dto.Status,
		Try_views:  dto.Try_views,
		UpdatedAt:  time.Now(),
	}
	c := ChaptersDao.Update(&Model, map[string]interface{}{
		"novel_id":   dto.Novel_id,
		"chapter_no": dto.Chapter_no,
		"title":      dto.Title,
		"desc":       dto.Desc,
		"link":       dto.Link,
		"is_vip":     dto.Is_vip,
		"source":     dto.Source,
		"views":      dto.Views,
		"text_num":   dto.Text_num,
		"status":     dto.Status,
		"try_views":  dto.Try_views,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

// Update
func (bs ChaptersService) UpdateStatus(dto dto.ChaptersUpdateStatusDto) int64 {
	Model := model.Chapters{
		Id: dto.Id,
	}
	c := ChaptersDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (bs ChaptersService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.Chapters{
		Id: dto.Id,
	}
	c := ChaptersDao.Delete(&Model)
	return c.RowsAffected
}

//上传封面
func (bs ChaptersService) UploadCover(file multipart.File, filename string) (filepath string, err error) {
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
	dst := "./data/images/" + strings.Replace(filename, ".", "_small.", 1)
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
