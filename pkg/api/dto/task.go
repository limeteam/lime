package dto

import (
	"lime/pkg/common"
	"time"
)

type TaskListDto struct {
	TaskName   string `form:"task_name" json:"task_name" `
	Status     string `form:"status" json:"status" `
	Time_type  string `form:"time_type,default=1" json:"time_type" `
	Start_time string `form:"start_time" json:"start_time" `
	End_time   string `form:"end_time" json:"end_time" `
	Page       string `form:"page" json:"page" `
	Skip       int    `form:"skip,default=0" json:"skip"`
	Limit      int    `form:"limit,default=20" json:"limit" binding:"max=100"`
}

type TaskCreateDto struct {
	TaskName          string            `form:"task_name" json:"task_name" `
	TaskRuleName      string            `form:"task_rule_name" json:"task_rule_name"`
	TaskDesc          string            `form:"task_desc" json:"task_desc"`
	Status            common.TaskStatus `form:"status" json:"status"`
	Counts            int               `form:"counts" json:"counts"`
	CronSpec          string            `form:"cron_spec" json:"cron_spec"`
	OutputType        string            `form:"output_type" json:"output_type"`
	OptUserAgent      string            `form:"opt_user_agent" json:"opt_user_agent"`
	OptMaxDepth       int               `form:"opt_max_depth" json:"opt_max_depth"`
	OptAllowedDomains string            `form:"opt_allowed_domains" json:"opt_allowed_domains"`
	OptURLFilters     string            `form:"opt_url_filters" json:"opt_url_filters"`
	OptMaxBodySize    int               `form:"opt_max_body_size" json:"opt_max_body_size"`
	OptRequestTimeout int               `form:"opt_request_timeout" json:"opt_request_timeout"`
	AutoMigrate       bool              `form:"auto_migrate" json:"auto_migrate"`
	LimitEnable       bool              `form:"limit_enable" json:"limit_enable"`
	LimitDomainRegexp string            `form:"limit_domain_regexp" json:"limit_domain_regexp" `
	LimitDomainGlob   string            `form:"limit_domain_glob" json:"limit_domain_glob" `
	LimitDelay        int               `form:"limit_delay" json:"limit_delay"`
	LimitRandomDelay  int               `form:"limit_random_delay" json:"limit_random_delay""`
	LimitParallelism  int               `form:"limit_parallelism" json:"limit_parallelism"`
	ProxyURLs         string            `form:"proxy_urls" json:"proxy_urls""`
	CreateTime        time.Time         `type(datetime)" json:"create_time"`
	LastLoginTime     time.Time         `type(datetime)" json:"-"`
}

type TaskEditDto struct {
	Id                int               `uri:"id" json:"id" binding:"required,gte=1"`
	TaskName          string            `form:"task_name" json:"task_name" `
	TaskRuleName      string            `form:"task_rule_name" json:"task_rule_name"`
	TaskDesc          string            `form:"task_desc" json:"task_desc"`
	Status            common.TaskStatus `form:"status" json:"status"`
	Counts            int               `form:"counts" json:"counts"`
	CronSpec          string            `form:"cron_spec" json:"cron_spec"`
	OutputType        string            `form:"output_type" json:"output_type"`
	OptUserAgent      string            `form:"opt_user_agent" json:"opt_user_agent"`
	OptMaxDepth       int               `form:"opt_max_depth" json:"opt_max_depth"`
	OptAllowedDomains string            `form:"opt_allowed_domains" json:"opt_allowed_domains"`
	OptURLFilters     string            `form:"opt_url_filters" json:"opt_url_filters"`
	OptMaxBodySize    int               `form:"opt_max_body_size" json:"opt_max_body_size"`
	OptRequestTimeout int               `form:"opt_request_timeout" json:"opt_request_timeout"`
	AutoMigrate       bool              `form:"auto_migrate" json:"auto_migrate"`
	LimitEnable       bool              `form:"limit_enable" json:"limit_enable"`
	LimitDomainRegexp string            `form:"limit_domain_regexp" json:"limit_domain_regexp" `
	LimitDomainGlob   string            `form:"limit_domain_glob" json:"limit_domain_glob" `
	LimitDelay        int               `form:"limit_delay" json:"limit_delay"`
	LimitRandomDelay  int               `form:"limit_random_delay" json:"limit_random_delay""`
	LimitParallelism  int               `form:"limit_parallelism" json:"limit_parallelism"`
	ProxyURLs         string            `form:"proxy_urls" json:"proxy_urls""`
}
