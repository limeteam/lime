package biquge

import (
	"lime/pkg/crawler/novels"
	"fmt"
)

//采集列表页面

func Test() {
	rules := novels.GetRules()
	fmt.Println(rules.RuleConfigInfo.GetSiteUrl.Pattern)
}
