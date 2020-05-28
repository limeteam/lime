package service

import (
	"errors"
	"lime/pkg/api/front/dao"
	"lime/pkg/api/front/dto"
)

var ChaptersDao = dao.ChaptersDao{}

type ChaptersService struct {}


func (bs ChaptersService) GetChapterInfoById(gdto dto.GeneralGetDto) (chapterInfo dto.ChapterInfoDto,err error) {
	chapterData := ChaptersDao.Get(gdto.Id)
	data := BooksDao.Get(chapterData.Novel_id)
	if data.Id < 1 {
		return chapterInfo,errors.New("未找到ID")
	}
	return dto.ChapterInfoDto{
		Book: dto.ChapterBookInfo{
			BookId:     data.Id,
			Name:       data.Name,
			IsFinished: data.Status,
		},
		ChapterTitle: chapterData.Title,
		Content:      chapterData.Desc,
		IsPreview:    0,
		LastChapter:  0,
		NextChapter:  "62759",
		ChapterId:    chapterData.Id,
		UpdateTime:   chapterData.UpdatedAt.String(),
	},nil
}

// List
func (bs ChaptersService) List(gdto dto.ChapterListDto) dto.ChaptersFrontListDto {
	data := BooksDao.Get(gdto.BookId)
	return dto.ChaptersFrontListDto{
		BookId:       data.Id,
		Name:         data.Name,
		Cover:        data.Cover,
		Description:  data.Desc,
		TotalChapter: data.Chapter_num,
		ChapterList:  getChapters(data.Id),
	}
}

func getChapters(bookId int) []dto.ChaptersFrontLists {
	chapters := ChaptersDao.GetListByBookId(bookId)
	var values []dto.ChaptersFrontLists
	for _, v := range chapters {
		chapter := dto.ChaptersFrontLists{
			BookId:v.Novel_id,
			ChapterId: v.Chapter_no,
			ChapterTitle: v.Title,
			Words: v.Text_num,
			IsVip: v.Is_vip,
			CanRead: v.Status,
			IsPreview:0,
			UpdateTime: v.UpdatedAt.String(),
			DisplayOrder: 0,
			LastChapter: 0,
			NextChapter: 2,
			Tag: dto.BookInfoTag{
				Tab:   "免费",
				Color: "#778899",
			},
		}
		values = append(values, chapter)
	}
	return values
}

