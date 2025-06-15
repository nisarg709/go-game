package models

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Tokens  int           `json:"tokens" bson:"tokens"`
	Tickets int           `json:"tickets" bson:"tickets"`
	Miles   int           `json:"miles" bson:"miles"`
	Badges  []string      `json:"badges" bson:"badges"`
}

func (u *User) HasTokens() bool {
	return u.Tokens > 0
}

func (u *User) HasEnoughTokens(num int) bool {
	return u.Tokens >= num
}
