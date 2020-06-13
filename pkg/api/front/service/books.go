package service

import (
	"fmt"
	"lime/pkg/api/admin/model"
	"lime/pkg/api/front/dao"
	"lime/pkg/api/front/dto"
	"math"
	"strconv"
	"time"
)

var BooksDao = dao.BooksDao{}
var CategoryDao = dao.CategoryDao{}

type BooksService struct{}


// List
func (bs BooksService) List(dto dto.GeneralListDto) ([]model.Books, int64) {
	return BooksDao.List(dto)
}


func (bs BooksService) GetBookInfoById(gdto dto.GeneralGetDto) (bookInfoDto dto.BookInfoDto, err error) {
	data := BooksDao.Get(gdto.Id)
	if data.Id < 1 {
		return bookInfoDto, nil
	}
	categories := CategoryDao.GetAll()
	categoryName := ""
	for categoryId, category := range categories {
		if data.Category_id == categoryId {
			categoryName = category.Name
		}
	}

	statusArr := map[int]string{0: "连载中", 1: "已完结", 2: "已太监"}
	statusName := statusArr[data.Status]
	views, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", data.Views/10000), 64)
	updateTime, _ := time.Parse("2006-01-02", data.Chapter_updated_at.String())
	TextNum := int(math.Floor(float64(data.Text_num / 10000)))
	var bookInfo = dto.BookInfoDto{
		BookId:          gdto.Id,
		Name:            data.Name,
		Cover:           data.Cover,
		Author:          data.Author,
		Description:     data.Desc,
		DisplayLabel:    fmt.Sprintf("%s·%s·%d万字", categoryName, statusName, TextNum),
		Finished:        statusName,
		Flag:            "热门",
		TotalWords:      fmt.Sprintf("%d万字", TextNum),
		TotalComment:    "45",
		ChapterLabel:    fmt.Sprintf("共%d章", data.Chapter_num),
		LastChapterTime: fmt.Sprintf("更新于%s", updateTime),
		LastChapter:     "第四百零七章 新的篇章",
		IsFinished:      data.Status,
		Score:           data.Score,
		Tags: []dto.BookInfoTag{
			{Tab: categoryName, Color: "#71c5fb"},
			{Tab: statusName, Color: "#f98445"},
		},
		Attribute: dto.BookAttribute{
			Popularity:      fmt.Sprintf("###%f###万人气", views),
			PopularityTitle: "人气飙升中",
			Reading:         fmt.Sprintf("###%f###万在读", views),
			ReadingTitle:    "在读人数攀升中",
			Score:           fmt.Sprintf("###%f###评分", data.Score),
			ScoreTitle:      "超过98%的同类书",
		},
		Labels: dto.BookInfoLabels{
			RecommendId: 0,
			Label:       "大家都在看",
			Style:       2,
			CanMore:     false,
			CanRefresh:  false,
			Total:       1,
			List:        getRandomBooks(gdto.Id),
		},
		IsCollect: 1,
	}
	return bookInfo, nil
}

func getRandomBooks(extraId int) []dto.BookInfoAllLookInfos {
	Books := BooksDao.GetRandBooks(extraId)
	categories := CategoryDao.GetAll()
	statusArr := map[int]string{0: "连载中", 1: "已完结", 2: "已太监"}
	var values []dto.BookInfoAllLookInfos
	for _, v := range Books {
		statusName := statusArr[v.Status]
		TextNum := int(math.Floor(float64(v.Text_num / 10000)))
		categoryName := ""
		for categoryId, category := range categories {
			if v.Category_id == categoryId {
				categoryName = category.Name
			}
		}
		book := dto.BookInfoAllLookInfos{
				BookId:      v.Id,
				Name:        v.Name,
				Cover:       v.Cover,
				Author:      v.Author,
				Description: v.Desc,
				Views:       v.Views,
				Tag: []dto.BookInfoTag{
					{Tab: categoryName, Color: "#71c5fb"},
					{Tab: statusName, Color: "#f98445"},
				},
				Finished:   statusName,
				Flag:       "",
				TotalWords: fmt.Sprintf("%d万字", TextNum),
				IsVip:      1,
				IsBaoyue:   1,
				IsFinished: v.Status,
		}
		values = append(values, book)
	}
	return values
}
