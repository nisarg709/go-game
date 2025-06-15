package models

import "github.com/globalsign/mgo/bson"

type Event struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string        `json:"title"`
	Body  string        `json:"body"`
	Image string        `json:"image"`
}
