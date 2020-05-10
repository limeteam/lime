package biquge

import (
	"fmt"
	"github.com/gocolly/colly"
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"lime/pkg/crawler/novels"
)

var ChapterService = service.ChapterService{}

func Chapter() {
	c := novels.Init()
	c.DetectCharset = true
	c.MaxDepth = 1
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url_prefix := e.Request.URL.Scheme + "://" + e.Request.URL.Host
		link := e.Attr("href")
		title := e.Text
		var chapterDto dto.ChapterCreateDto
		chapterDto.Url = url_prefix + link
		chapterDto.Book_id = 3
		chapterDto.Title = title
		created := ChapterService.Create(chapterDto)
		if created.Id <= 0 {
			fmt.Println("error add")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://www.xbaquge.com/files/article/html/40/40670/")
}
