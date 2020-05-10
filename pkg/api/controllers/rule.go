package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"lime/pkg/crawler/core"
)

type RuleController struct {
	BaseController
}

func (r *RuleController) GetRuleList(c *gin.Context) {
	keys := core.GetTaskRuleKeys()
	if len(keys) == 0 {
		log.Warnf("task rule is empty")
	}
	resp(c, map[string]interface{}{
		"list":  keys,
		"total": len(keys),
	})
}
