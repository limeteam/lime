package novels

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/redisstorage"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	//"github.com/gocolly/colly/debug"
	"log"
)

type Rules struct {
	RuleConfigInfo RuleconfigInfo `yaml:"RuleConfigInfo"`
}

type Rule struct {
	FilterPattern string `yaml:"FilterPattern"`
	Method        string `yaml:"Method"`
	Options       string `yaml:"Options"`
	Pattern       string `yaml:"Pattern"`
	RegexName     string `yaml:"RegexName"`
}
type RuleconfigInfo struct {
	GetSiteCharset           Rule `yaml:"GetSiteCharset"`
	GetSiteName              Rule `yaml:"GetSiteName"`
	GetSiteUrl               Rule `yaml:"GetSiteUrl"`
	LagerSort                Rule `yaml:"LagerSort"`
	NovelAuthor              Rule `yaml:"NovelAuthor"`
	NovelCover               Rule `yaml:"NovelCover"`
	NovelDefaultCoverUrl     Rule `yaml:"NovelDefaultCoverUrl"`
	NovelDegree              Rule `yaml:"NovelDegree"`
	NovelErr                 Rule `yaml:"NovelErr"`
	NovelInfo_GetNovelPubKey Rule `yaml:"NovelInfo_GetNovelPubKey"`
	NovelIntro               Rule `yaml:"NovelIntro"`
	NovelKeyword             Rule `yaml:"NovelKeyword"`
	NovelList_GetNovelKey    Rule `yaml:"NovelList_GetNovelKey"`
	NovelListUrl             Rule `yaml:"NovelListUrl"`
	NovelName                Rule `yaml:"NovelName"`
	NovelSearch_GetNovelKey  Rule `yaml:"NovelSearch_GetNovelKey"`
	NovelSearchData          Rule `yaml:"NovelSearchData"`
	NovelSearchUrl           Rule `yaml:"NovelSearchUrl"`
	NovelUrl                 Rule `yaml:"NovelUrl"`
	PubChapter_GetChapterKey Rule `yaml:"PubChapter_GetChapterKey"`
	PubChapterName           Rule `yaml:"PubChapterName"`
	PubContent_GetTextKey    Rule `yaml:"PubContent_GetTextKey"`
	PubContentErr            Rule `yaml:"PubContentErr"`
	PubContentImages         Rule `yaml:"PubContentImages"`
	PubContentPage           Rule `yaml:"PubContentPage"`
	PubContentReplace        Rule `yaml:"PubContentReplace"`
	PubContentText           Rule `yaml:"PubContentText"`
	PubContentTitle          Rule `yaml:"PubContentTitle"`
	PubContentUrl            Rule `yaml:"PubContentUrl"`
	PubCookies               Rule `yaml:"PubCookies"`
	PubIndexErr              Rule `yaml:"PubIndexErr"`
	PubIndexUrl              Rule `yaml:"PubIndexUrl"`
	PubTextUrl               Rule `yaml:"PubTextUrl"`
	PubVolumeContent         Rule `yaml:"PubVolumeContent"`
	PubVolumeName            Rule `yaml:"PubVolumeName"`
	PubVolumeSplit           Rule `yaml:"PubVolumeSplit"`
	RuleID                   Rule `yaml:"RuleID"`
	RuleVersion              Rule `yaml:"RuleVersion"`
	SmallSort                Rule `yaml:"SmallSort"`
	TsContrary               Rule `yaml:"TsContrary"`
}

func Init() *colly.Collector {
	storage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "test",
	}
	c := colly.NewCollector(
	//colly.AllowedDomains("android.myapp.com"),
	//colly.Debugger(&debug.LogDebugger{}),
	)
	err := c.SetStorage(storage)
	if err != nil {
		panic(err)
	}

	// delete previous data from storage
	if err := storage.Clear(); err != nil {
		log.Fatal(err)
	}

	// close redis client
	//defer storage.Client.Close()
	return c
}

func GetRules() *Rules {
	var rules Rules
	config, err := ioutil.ReadFile("./config/rules/biquge.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &rules)
	return &rules
}
