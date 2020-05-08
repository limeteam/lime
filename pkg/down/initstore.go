package down

import (
	"lime/pkg/down/site"
	"lime/pkg/down/store"
	"fmt"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"time"
)

var chapter *store.Store

func InitLoadStore(driver string, bookurl string) error {
	if bookurl == "" {
		return fmt.Errorf("请输入小说的链接！")
	}
	bookURL, err := url.Parse(bookurl)
	if err != nil {
		return err
	}
	log.Printf("URL: %#v", bookURL.String())
	switch driver {
	case "phantomjs":
		log.Printf("初始化 PhantomJS")
		site.InitPhantomJS()
		defer func() {
			log.Printf("关闭 PhantomJS")
			site.ClosePhantomJS()
		}()
		for errCount := 0; errCount < 20; errCount++ {
			chapter, err = site.PhBookInfo(bookURL.String())
			if err == nil {
				break
			} else {
				log.Printf("ErrCount: %d Err: %s", errCount, err)
				time.Sleep(1000 * time.Millisecond)
			}
		}
	case "chromedp":
		log.Printf("Chromedp Running...")
		chapter, err = site.ChromedpBookInfo(bookURL.String(), "chromedp-log")
	default:
		log.Printf("使用http方式下载")
		for errCount := 0; errCount < 20; errCount++ { //重试20次
			chapter, err = site.BookInfo(bookURL.String())
			if err == nil {
				break
			} else {
				log.Printf("ErrCount: %d Err: %s", errCount, err)
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
	if err != nil {
		return err
	}
	chapter.BookURL = bookURL.String()
	filename := fmt.Sprintf("%s.%s", chapter.BookName, store.FileExt)
	filemode, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		bookContent, err := yaml.Marshal(chapter)
		if err != nil {
			return err
		}
		ioutil.WriteFile(filename, bookContent, 0775)
		return nil
	}
	if filemode.IsDir() {
		return fmt.Errorf("is Dir")
	}
	log.Printf("Loading....")
	bookContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bookContent, &chapter)
	if err != nil {
		return err
	}
	return nil
}

func LoadLocalStore(filename string) error {
	log.Printf("Loading cache file: %s", filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &chapter)
	if err != nil {
		return err
	}
	return nil
}
