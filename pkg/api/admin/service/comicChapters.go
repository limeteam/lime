package service

import (
	"fmt"
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

var ComicChaptersDao = dao.ComicChaptersDao{}

// Service
type ComicChaptersService struct {
}

// InfoOfId
func (bs ComicChaptersService) InfoOfId(dto dto.GeneralGetDto) model.ComicChapters {
	return ComicChaptersDao.Get(dto.Id)
}

func (bs ComicChaptersService) GetAll() []model.ComicChapters {
	return ComicChaptersDao.GetAll()
}

// List
func (bs ComicChaptersService) List(dto dto.GeneralListDto) ([]model.ComicChapters, int64) {
	return ComicChaptersDao.List(dto)
}

// Create
func (bs ComicChaptersService) Create(dto dto.ComicChaptersCreateDto) (model.ComicChapters, error) {
	Model := model.ComicChapters{
		Comic_id:    dto.Comic_id,
		Chapter_no:  dto.Chapter_no,
		Title:       dto.Title,
		Desc:        dto.Desc,
		Cover:       dto.Cover,
		Upload_type: dto.Upload_type,
		Is_vip:      dto.Is_vip,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	c := ComicChaptersDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (bs ComicChaptersService) Update(dto dto.ComicChaptersEditDto) int64 {
	Model := model.ComicChapters{
		Id: dto.Id,
	}
	fmt.Println(dto.Desc)
	var desc = dto.Desc
	c := ComicChaptersDao.Update(&Model, map[string]interface{}{
		"comic_id":    dto.Comic_id,
		"chapter_no":  dto.Chapter_no,
		"title":       dto.Title,
		"desc":        desc,
		"cover":       dto.Cover,
		"upload_type": dto.Upload_type,
		"is_vip":      dto.Is_vip,
		"updated_at":  time.Now(),
	})
	return c.RowsAffected
}

// Update
func (bs ComicChaptersService) UpdateStatus(dto dto.ComicChaptersUpdateStatusDto) int64 {
	Model := model.ComicChapters{
		Id: dto.Id,
	}
	c := ComicChaptersDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (bs ComicChaptersService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.ComicChapters{
		Id: dto.Id,
	}
	c := ComicChaptersDao.Delete(&Model)
	return c.RowsAffected
}

//上传封面
func (bs ComicChaptersService) UploadCover(file multipart.File, filename string) (filepath string, err error) {
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
