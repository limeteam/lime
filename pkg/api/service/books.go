package service

import (
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"io"
	"lime/pkg/api/dao"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
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
		UpdatedAt:                 time.Now(),
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

func (bs BooksService) GetBookInfoById(gdto dto.GeneralGetDto) dto.BookInfoDto {
	var bookinfo = dto.BookInfoDto{
		BookId: gdto.Id,
		Name:   "test",
		Cover: "http://beiwo-new.oss-cn-beijing.aliyuncs.com/cover/84/7757bb60adec76a2166f8af2cde68117.jpeg?x-oss-process=image%2Fresize%2Cw_300%2Ch_400%2Cm_lfit",
		Author: "孤梦夕阳",
		Description: "这是一个充满元素气息的世界—鸿界。各种元素游荡在世界的各个角落，在这环境下孕育而生的生命经过漫长时间的演化，生活，学习，争夺，繁衍下对于各种元素的掌握到达了一个巅峰，迎来一个大时代—三极天下。鸿界分为...",
		DisplayLabel: "网游竞技·连载中·141万字",
		Finished: "连载中",
		Flag: "热门",
		TotalWords: "141万字",
		TotalComment: "45",
		ChapterLabel: "共407章",
		LastChapterTime: "更新于2020-02-02",
		LastChapter: "第四百零七章 新的篇章",
		IsFinished: 1,
		Score: "",
		Tags: []dto.BookInfoTag{
			{Tab: "网游竞技", Color: "#71c5fb"},
			{Tab: "已完结", Color: "#f98445"},
		},
		Attribute: dto.BookAttribute{
			Popularity:      "###1.0###万人气",
			PopularityTitle: "人气飙升中",
			Reading:         "###1.0###万在读",
			ReadingTitle:    "在读人数攀升中",
			Score:           "###7.0###评分",
			ScoreTitle:      "超过98%的同类书",
		},
		Labels: dto.BookInfoLabels{
			RecommendId: 0,
			Label:       "大家都在看",
			Style:       2,
			CanMore:     false,
			CanRefresh:  false,
			Total:       1,
			List: []dto.BookInfoAllLookInfos{
				{
					BookId:      2,
					Name:        "阴阳诡路",
					Cover:       "http://beiwo-new.oss-cn-beijing.aliyuncs.com/cover/212/f73c9adb1f5e069cbeee5c7f600dfd1c.jpeg?x-oss-process=image%2Fresize%2Cw_300%2Ch_400%2Cm_lfit",
					Author:      "牧雪",
					Description: "有一种人，能去祸免灾、驱邪避凶、普怨度灵、识破天机、断人生死、游走于阴阳两界。一个接一个的邪灵恶煞会出现于他们的面前，一件又一件的奇异诡事会发生在他们的身边，于是这种人通常被称作：阴阳先生。我是阴阳女...",
					Views:       10,
					Tag: []dto.BookInfoTag{
						{Tab: "网游竞技", Color: "#71c5fb"},
						{Tab: "已完结", Color: "#f98445"},
					},
					Finished:   "已完结",
					Flag:       "",
					TotalWords: "55万字",
					IsVip:      1,
					IsBaoyue:   1,
					IsFinished: 1,
				},
			},
		},
		IsCollect: 1,
	}

	return bookinfo
}
