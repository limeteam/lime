package aoyuge

import (
	"github.com/axgle/mahonia"
	log "github.com/sirupsen/logrus"
	"lime/pkg/crawler/core"
)

func init() {
	core.Register(rule)
}

var rule = &core.TaskRule{
	Name:         "傲宇阁小说爬虫规则",
	Description:  "下载傲宇阁小说内容",
	Namespace:    "aoyuge",
	OutputFields: []string{"name", "author", "image", "url", "desc"},
	Rule: &core.Rule{
		Head: func(ctx *core.Context) error {
			return ctx.VisitForNext("https://www.aoyuge.com/31/31942/")
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
		"html": func(ctx *core.Context, el *core.HTMLElement) error {
			name := el.ChildAttr(`meta[property="og:novel:book_name"]`, "content")
			name = ConvertToString(name, "gbk", "utf-8")
			author := el.ChildAttr(`meta[property="og:novel:author"]`, "content")
			author = ConvertToString(author, "gbk", "utf-8")
			image := el.ChildAttr(`meta[property="og:image"]`, "content")
			url := el.ChildAttr(`meta[property="og:novel:read_url"]`, "content")
			desc := el.ChildAttr(`meta[property="og:description"]`, "content")
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
