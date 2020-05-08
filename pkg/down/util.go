package down

import (
	"regexp"
)

func TitleAlias(s string) (alias []string) {
	var res = []string{
		`^([第]?([零一二三四五六七八就十百千万0-9]+)[章\,，、]?)[ ]*([^ （）\(\))]+).*$`,
		`([^ （）\(\)]+)`,   // 没有章节号的章节
		`[（\(](.+?)[）\)]`, //括号内的
	}
	for _, v := range res {
		re, err := regexp.Compile(v)
		if err != nil {
			panic(err)
		}
		find := re.FindAllStringSubmatch(s, -1)
		if find == nil {
			continue
		}
		for _, v1 := range find {
			for _, v := range v1[1:] {
				if v == "" ||  v == s || StringInSlice(v, alias){
					continue
				}
				alias = append(alias, v)
			}
		}
	}
	return
}

func StringInSlice(str string, slice []string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}