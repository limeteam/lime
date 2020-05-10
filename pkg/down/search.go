package down

import (
	"fmt"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"lime/pkg/down/site"
)

func Search(keyword string, put bool, filename string, driver string) error {
	var bookurl = ""
	r, err := site.Search(keyword)
	if err != nil {
		return err
	}
	if !put {
		fmt.Printf("搜索到%d个内容:\n", len(r))
		for _, v := range r {
			fmt.Printf("%s %s %s\n", v.BookURL, v.BookName, v.Author)
		}
		return nil
	}
	if filename == "" {
		err := InitLoadStore(driver, bookurl)
		if err != nil {
			return err
		}
	} else {
		err := LoadLocalStore(filename)
		if err != nil {
			return err
		}
	}
	rrr := []site.ChaperSearchResult{}
	for _, v := range r {
		if (v.Author == chapter.Author) && (v.BookName == chapter.BookName) {
			log.Printf("%s %s %s", v.BookURL, v.BookName, v.Author)
			rrr = append(rrr, v)
		}
	}
	b, err := yaml.Marshal(chapter)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filename, b, 0775)
	return nil
}
