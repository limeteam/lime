package service

import (
	"fmt"
	"lime/pkg/api/front/dao"
	"lime/pkg/api/front/dto"
	"math"
	"strconv"
	"time"
)

var BooksDao = dao.BooksDao{}
var CategoryDao = dao.CategoryDao{}

type BooksService struct {}

func (bs BooksService) GetBookInfoById(gdto dto.GeneralGetDto) (bookInfoDto dto.BookInfoDto, err error) {
	data := BooksDao.Get(gdto.Id)
	if data.Id < 1 {
		return  bookInfoDto,nil
	}
	categories := CategoryDao.GetAll()
	categoryName := ""
	for categoryId,category := range categories {
		if data.Category_id == categoryId {
			categoryName = category.Name
		}
	}

	statusArr := map[int]string{0: "连载中",1: "已完结",2: "已太监"}
	statusName := statusArr[data.Status]
	views, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", data.Views/10000), 64)
	fmt.Println(data.Chapter_updated_at)
	updateTime,_ := time.Parse("2006-01-02",data.Chapter_updated_at.String())
	TextNum := int(math.Floor(float64(data.Text_num/10000)))
	var bookInfo = dto.BookInfoDto{
		BookId:          gdto.Id,
		Name:            data.Name,
		Cover:           data.Cover,
		Author:          data.Author,
		Description:     data.Desc,
		DisplayLabel:    fmt.Sprintf("%s·%s·%d万字",categoryName,statusName,TextNum),
		Finished:        statusName,
		Flag:            "热门",
		TotalWords:      fmt.Sprintf("%d万字",TextNum),
		TotalComment:    "45",
		ChapterLabel:    fmt.Sprintf("共%d章",data.Chapter_num),
		LastChapterTime: fmt.Sprintf("更新于%s", updateTime),
		LastChapter:     "第四百零七章 新的篇章",
		IsFinished:      data.Status,
		Score:           data.Score,
		Tags: []dto.BookInfoTag{
			{Tab: categoryName, Color: "#71c5fb"},
			{Tab: statusName, Color: "#f98445"},
		},
		Attribute: dto.BookAttribute{
			Popularity:      fmt.Sprintf("###%f###万人气",views),
			PopularityTitle: "人气飙升中",
			Reading:         fmt.Sprintf("###%f###万在读",views),
			ReadingTitle:    "在读人数攀升中",
			Score:           fmt.Sprintf("###%f###评分",data.Score),
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
	return bookInfo,nil
}
