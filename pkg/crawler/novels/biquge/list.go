package biquge

import (
	"fmt"
	"lime/pkg/crawler/novels"
)

//采集列表页面

func Test() {
	rules := novels.GetRules()
	fmt.Println(rules.RuleConfigInfo.GetSiteUrl.Pattern)
}
