package aoyuge

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"net/url"
	"os"
	"sync"
)

// Chapter 章节
type Chapter struct {
	Name    string //名称
	URL     string //章节链接
	Text    []string
	Alias   []string   `yaml:"-"`
	MuxLock sync.Mutex `yaml:"-"`
}

const FileExt = "lime"

// Volume 卷
type Volume struct {
	Name     string
	IsVIP    bool
	Chapters []Chapter
}

// Store is store yaml data file format
type Store struct {
	BookURL     string
	BookName    string
	Author      string // 作者
	CoverURL    string // 封面链接
	Description string // 介绍
	Volumes     []Volume
}

func WriteBook(bookurl string, chapter Store) error {
	bookURL, err := url.Parse(bookurl)
	if err != nil {
		return err
	}
	chapter.BookURL = bookURL.String()
	filename := fmt.Sprintf("%s.%s", chapter.BookName, FileExt)
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
	return nil
}
