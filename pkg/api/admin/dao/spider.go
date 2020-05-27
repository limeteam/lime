package dao

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
	"time"
)

type SpiderDao struct{}

func (sd SpiderDao) Add(row map[int]interface{}) error {
	var spider model.Spider
	s, c := db.GetCol("spiders")
	defer s.Close()
	spider.Id = bson.NewObjectId()
	spider.AppName = fmt.Sprintf("%v", row[0])

	spider.CreatedAt = time.Now()
	spider.UpdatedAt = time.Now()
	if err := c.Insert(&spider); err != nil {
		return err
	}
	return nil
}

func (sd SpiderDao) InitYaml(row map[int]interface{}) error {

	return nil
}

func (sd SpiderDao) AddYaml(row map[int]interface{}) error {

	return nil
}
