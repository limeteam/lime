package store

import "sync"

const FileExt = "lime"

// Volume 卷
type Volume struct {
	Name     string
	IsVIP    bool
	Chapters []Chapter
}

// Chapter 章节
type Chapter struct {
	Name    string //名称
	URL     string //章节链接
	Text    []string
	Alias   []string   `yaml:"-"`
	MuxLock sync.Mutex `yaml:"-"`
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

func (store Store) Total() (Done, AllChaper int) {
	for _, v := range store.Volumes {
		AllChaper += len(v.Chapters)
		for _, v2 := range v.Chapters {
			if len(v2.Text) != 0 {
				Done++
			}
		}
	}
	return
}
