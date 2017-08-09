package models

import (
	"labix.org/v2/mgo/bson"
)

type UserCredentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Age      int           `json:"age" bson:"age"`
	Gender   string        `json:"gender" bson:"gender"`
	Type     string        `json:"type" bson:"type"`
	Email    string        `json:"email" bson:"email"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}
