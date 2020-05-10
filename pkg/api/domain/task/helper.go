package task

import (
	"github.com/pkg/errors"
	"lime/pkg/api/model"
	"lime/pkg/crawler/core"
	"regexp"
	"strings"
	"time"
)

func GetSpiderTaskByModel(task model.Task) (*core.Task, error) {
	rule, err := core.GetTaskRule(task.TaskRuleName)
	if err != nil {
		return nil, err
	}

	var optAllowedDomains []string
	if task.OptAllowedDomains != "" {
		optAllowedDomains = strings.Split(task.OptAllowedDomains, ",")
	}
	var urlFiltersReg []*regexp.Regexp
	if task.OptURLFilters != "" {
		urlFilters := strings.Split(task.OptURLFilters, ",")
		for _, v := range urlFilters {
			reg, err := regexp.Compile(v)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			urlFiltersReg = append(urlFiltersReg, reg)
		}
	}

	config := core.TaskConfig{
		CronSpec: task.CronSpec,
		Option: core.Option{
			UserAgent:              task.OptUserAgent,
			MaxDepth:               task.OptMaxDepth,
			AllowedDomains:         optAllowedDomains,
			URLFilters:             urlFiltersReg,
			AllowURLRevisit:        rule.AllowURLRevisit,
			MaxBodySize:            task.OptMaxBodySize,
			IgnoreRobotsTxt:        rule.IgnoreRobotsTxt,
			InsecureSkipVerify:     rule.InsecureSkipVerify,
			ParseHTTPErrorResponse: rule.ParseHTTPErrorResponse,
			DisableCookies:         rule.DisableCookies,
		},
		Limit: core.Limit{
			Enable:      task.LimitEnable,
			DomainGlob:  task.LimitDomainGlob,
			Delay:       time.Duration(task.LimitDelay) * time.Millisecond,
			RandomDelay: time.Duration(task.LimitRandomDelay) * time.Millisecond,
			Parallelism: task.LimitParallelism,
		},
		OutputConfig: core.OutputConfig{
			Type: task.OutputType,
		},
	}

	if task.OptRequestTimeout > 0 {
		config.Option.RequestTimeout = time.Duration(task.OptRequestTimeout) * time.Second
	}
	if urls := strings.TrimSpace(task.ProxyURLs); len(urls) > 0 {
		config.ProxyURLs = strings.Split(urls, ",")
	}

	return core.NewTask(task.Id, *rule, config), nil
}
