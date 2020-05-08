package down

import (
	"lime/pkg/down/store"
	"fmt"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	processbar "gopkg.in/cheggaaa/pb.v1"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
)

var threadsNum = 10

func Download(bookurl string, driver string) error {
	var filename = ""
	if err := InitLoadStore(driver, bookurl); err != nil {
		return err
	}
	//var err error
	printChapterInfo(chapter)//打印小说信息
	SyncStore := &SyncStore{
		Store: chapter,
	}
	SyncStore.Init()
	chCount,isDone := syncStoreChapter(chapter) //载入小说
	if isDone < chCount {
		bar := processbar.StartNew(chCount)
		bar.Set(isDone)
		Jobch := make(chan error)
		for i := 0; i < threadsNum; i++ {
			go Job(SyncStore, Jobch)
		}
		cc := make(chan os.Signal)
		signal.Notify(cc, os.Interrupt)

		var ii = 0
	LoopStep1:
		for {
			select {
			case ccc := <-cc:
				log.Printf("进程信号: %v", ccc)
				return nil
			case err := <-Jobch:
				if err != nil {
					if err == io.EOF {
						ii++
						if ii >= threadsNum {
							bar.Finish()
							log.Printf("缓存完成")
							break LoopStep1
						}
					} else {
						log.Printf("Job Error: %s", err)
					}
				} else {
					bar.Increment()
				}
			}
		}
		close(Jobch)
		log.Printf("生成别名")
		for k3, vol2 := range chapter.Volumes {
			for k4, chaper := range vol2.Chapters {
				chapter.Volumes[k3].Chapters[k4].Alias = TitleAlias(chaper.Name)
			}
		}
	}
	chapter = replaceChapterString(chapter) //替换特殊内容
	return writeFile(filename,chapter)
}

func printChapterInfo(chapter *store.Store) {
	fmt.Printf("书名: %#v\n", chapter.BookName)
	fmt.Printf("作者: %#v\n", chapter.Author)
	fmt.Printf("封面: %s\n", chapter.CoverURL)
	fmt.Printf("简介: \n\t%v\n", strings.Replace(chapter.Description, "\n", "\n\t", -1))
	fmt.Printf("章节数: \n")
	for _, v := range chapter.Volumes {
		var VIP string
		if v.IsVIP {
			VIP = "VIP"
		} else {
			VIP = "免费"
		}
		fmt.Printf("\t%s卷(%s) %d章\n", v.Name, VIP, len(v.Chapters))
	}
	log.Printf("线程数: %d,预缓存中...\n", threadsNum)
}

func syncStoreChapter(chapter *store.Store) (chCount int,isDone int){
	chCount = 0
	isDone = 0
	for _, v := range chapter.Volumes {
		chCount += len(v.Chapters)
		for _, chapter := range v.Chapters {
			if len(chapter.Text) != 0 {
				isDone++
			}
		}
	}
	if isDone != 0 {
		log.Printf("[读入] 已缓存:%d", isDone)
	}
	// End Print
	defer func(s *store.Store) {
		var chCount = 0
		var isDone = 0
		for _, v := range chapter.Volumes {
			chCount += len(v.Chapters)
			for _, v2 := range v.Chapters {
				if len(v2.Text) != 0 {
					isDone++
				}
			}
		}
		if isDone != 0 {
			log.Printf("[爬取结束] 已缓存:%d", isDone)
		}
	}(chapter)
	return
}

//写入文件
func writeFile(filename string,chapter *store.Store) error {
	b, err := yaml.Marshal(chapter)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, b, 0775); err != nil {
		return err
	}
	return nil
}

func replaceChapterString(chapter *store.Store) (chapters *store.Store) {
	for k3,Volume  := range chapter.Volumes {
		for k4, ch2 := range Volume.Chapters {
			newContent := []string{}
			for _, v := range ch2.Text {
				v = strings.TrimSpace(v)
				if v == "" {
					continue
				}
				v = strings.Replace(v, "“”", "", -1)
				if regexp.MustCompile("^[…]+$").MatchString(v) {
					continue
				}
				newContent = append(newContent, v)
			}
			chapter.Volumes[k3].Chapters[k4].Text = newContent
		}
	}
	return chapter
}