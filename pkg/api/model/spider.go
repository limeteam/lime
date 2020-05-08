package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Spider struct {
	Id bson.ObjectId `json:"_id" bson:"_id"`
	AppName string 	`json:"app_name" bson:"app_name"`
	CreatedAt time.Time `json:"create_at" bson:"create_at"`
	UpdatedAt time.Time `json:"update_at" bson:"update_at"`
}

