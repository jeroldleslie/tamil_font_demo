package models

import (
	"labix.org/v2/mgo/bson"
)

type Question struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Question     string        `json:"question" bson:"question"`
}
