package fanfan

import (
	"github.com/axgle/mahonia"
	log "github.com/sirupsen/logrus"
	"lime/pkg/crawler/core"
)

func init() {
	core.Register(rule)
}

var rule = &core.TaskRule{
	Name:         "饭饭小说爬虫规则",
	Description:  "下载饭饭小说内容",
	Namespace:    "ffxs",
	OutputFields: []string{"name", "author", "image", "url", "desc"},
	Rule: &core.Rule{
		Head: func(ctx *core.Context) error {
			return ctx.VisitForNext("http://ffxs.me/dsyq/11605")
		},
		Nodes: map[int]*core.Node{
			0: step1, // 第一步: 获取小说简介
		},
	},
}

var step1 = &core.Node{
	OnError: func(ctx *core.Context, res *core.Response, err error) error {
		log.Errorf("Visiting failed! url:%s, err:%s", res.Request.URL.String(), err.Error())
		return nil
	},
	OnRequest: func(ctx *core.Context, req *core.Request) {
		log.Println("Visting", req.URL.String())
	},
	OnHTML: map[string]func(*core.Context, *core.HTMLElement) error{
		"div[class=info-right]": func(ctx *core.Context, el *core.HTMLElement) error {
			name := el.ChildText("div[class=desc] > h1")
			name = ConvertToString(name, "gbk", "utf-8")
			author := el.DOM.Find("div[class=desc]").Eq(1).Text()
			author = ConvertToString(author, "gbk", "utf-8")
			image := el.ChildAttr("div[class=info-left] div[class=cover] > img", "src")
			url := el.Request.URL.String()
			desc := el.ChildText("div[class=descInfo] > p")
			desc = ConvertToString(desc, "gbk", "utf-8")
			return ctx.Output(map[int]interface{}{
				0: name,
				1: author,
				2: image,
				3: url,
				4: desc,
			})
		},
	},
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
