package biquge

import (
	"lime/pkg/crawler/novels"
	"fmt"
	"github.com/gocolly/colly"
	//"regexp"
)

var rules = novels.GetRules().RuleConfigInfo

func Book() {
	c := novels.Init()
	c.DetectCharset = true
	c.MaxDepth = 1
	c.OnHTML("html", func(e *colly.HTMLElement) {
		//intro,_ := e.DOM.Find("#maininfo").Html()
		//re, _ := regexp.Compile(rules.NovelIntro.Pattern)
		//NovelIntro := re.FindString(intro)
		//fmt.Println(NovelIntro)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://www.xbaquge.com/files/article/html/40/40670/")
}
