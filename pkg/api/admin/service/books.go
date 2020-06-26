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

var BooksDao = dao.BooksDao{}

// Service
type BooksService struct {
}

// InfoOfId
func (bs BooksService) InfoOfId(dto dto.GeneralGetDto) model.Books {
	return BooksDao.Get(dto.Id)
}

func (bs BooksService) GetAll() []model.Books {
	return BooksDao.GetAll()
}

// List
func (bs BooksService) List(dto dto.GeneralListDto) ([]model.Books, int64) {
	return BooksDao.List(dto)
}

// Create
func (bs BooksService) Create(dto dto.BooksCreateDto) (model.Books, error) {
	Model := model.Books{
		Name:                      dto.Name,
		Old_name:                  dto.Old_name,
		Channel_id:                dto.Channel_id,
		Category_id:               dto.Category_id,
		Desc:                      dto.Desc,
		Cover:                     dto.Cover,
		Author:                    dto.Author,
		Source:                    dto.Source,
		Split_rule:                dto.Split_rule,
		Upload_file:               dto.Upload_file,
		Status:                    dto.Status,
		Flag:                      dto.Flag,
		Chapter_price:             dto.Chapter_price,
		Thousand_characters_price: dto.Thousand_characters_price,
		Score:                     dto.Score,
		Text_num:                  dto.Text_num,
		Chapter_num:               dto.Chapter_num,
		Chapter_updated_at:        dto.Chapter_updated_at,
		Chapter_id:                dto.Chapter_id,
		Chapter_title:             dto.Chapter_title,
		Views:                     dto.Views,
		Collect_num:               dto.Collect_num,
		Online_status:             dto.Online_status,
		Is_sensitivity:            dto.Is_sensitivity,
		CreatedAt:                 time.Now(),
		UpdatedAt:                 time.Now(),
	}
	c := BooksDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (bs BooksService) Update(dto dto.BooksEditDto) int64 {
	Model := model.Books{
		Id:                        dto.Id,
	}
	c := BooksDao.Update(&Model, map[string]interface{}{
		"name":                      dto.Name,
		"old_name":                  dto.Old_name,
		"channel_id":                dto.Channel_id,
		"category_id":               dto.Category_id,
		"desc":                      dto.Desc,
		"cover":                     dto.Cover,
		"author":                    dto.Author,
		"source":                    dto.Source,
		"split_rule":                dto.Split_rule,
		"upload_file":               dto.Upload_file,
		"status":                    dto.Status,
		"flag":                      dto.Flag,
		"chapter_price":             dto.Chapter_price,
		"thousand_characters_price": dto.Thousand_characters_price,
		"score":                     dto.Score,
		"text_num":                  dto.Text_num,
		"chapter_num":               dto.Chapter_num,
		"chapter_updated_at":        dto.Chapter_updated_at,
		"chapter_id":                dto.Chapter_id,
		"chapter_title":             dto.Chapter_title,
		"views":                     dto.Views,
		"collect_num":               dto.Collect_num,
		"online_status":             dto.Online_status,
		"is_sensitivity":            dto.Is_sensitivity,
		"updated_at":                time.Now(),
	})
	return c.RowsAffected
}

// Update
func (bs BooksService) UpdateStatus(dto dto.BooksUpdateStatusDto) int64 {
	Model := model.Books{
		Id: dto.Id,
	}
	c := BooksDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (bs BooksService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.Books{
		Id: dto.Id,
	}
	c := BooksDao.Delete(&Model)
	return c.RowsAffected
}

//上传封面
func (bs BooksService) UploadCover(file multipart.File, filename string) (filepath string, err error) {
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
func (bs BooksService) UploadBookFile(file multipart.File, filename string) (filepath string, err error) {
	defer file.Close()
	filenameWithSuffix := path.Base(filename)
	fileSuffix := path.Ext(filenameWithSuffix)
	filename = strconv.FormatInt(time.Now().Unix(), 10) + fileSuffix
	destFile, err := os.Create("./data/books/" + filename)
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
	filepath = "/upload/books/" + filename
	return filepath, nil
}

