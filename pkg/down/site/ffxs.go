package site

import (
	"lime/pkg/down/store"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net/url"
	"strings"
)

var ffxs = SiteA{
	Name:     "饭饭小说",
	HomePage: "https://www.ffxs.me/",
	Match: []string{
		`https://www\.ffxs\.me/[a-z]+/\d+/*`,
		`https://www\.ffxs\.me/book/\d+-\d+-\d+\.html`,
	},
	BookInfo: func(body io.Reader) (s *store.Store, err error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return
		}
		s = &store.Store{}
		s.BookName = htmlquery.InnerText(htmlquery.FindOne(doc, `//div[@class="desc"]/h1`))
		node_content := htmlquery.Find(doc, `//*[@class="catalog"]/ul/li/a`)
		if len(node_content) == 0 {
			err = fmt.Errorf("No matching contents")
			return
		}
		var vol = store.Volume{
			Name:     "正文",
			Chapters: make([]store.Chapter, 0),
		}

		for _, v := range node_content {
			chapterURL, err := url.Parse(htmlquery.SelectAttr(v, "href"))
			if err != nil {
				return nil, err
			}

			vol.Chapters = append(vol.Chapters, store.Chapter{
				Name: strings.TrimSpace(htmlquery.InnerText(v)),
				URL:  chapterURL.String(),
			})
		}
		s.Volumes = append(s.Volumes, vol)

		return
	},
	Chapter: func(body io.Reader) ([]string, error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return nil, err
		}

		M := []string{}
		//list
		nodeContent := htmlquery.Find(doc, `//*[@class="content"]/text()`)
		if len(nodeContent) == 0 {
			err = fmt.Errorf("No matching content")
			return nil, err
		}
		for _, v := range nodeContent {
			t := htmlquery.InnerText(v)
			t = strings.TrimSpace(t)

			if strings.HasPrefix(t, "…") {
				continue
			}

			t = strings.Replace(t, "…", "", -1)
			t = strings.Replace(t, ".......", "", -1)
			t = strings.Replace(t, "...”", "”", -1)

			if t == "" {
				continue
			}

			M = append(M, t)
		}
		return M, nil
	},
}
