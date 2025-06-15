package models

import "github.com/globalsign/mgo/bson"

type OddOneOutQuestion struct {
	Id           bson.ObjectId `bson:"_id"`
	CategoryType string        `bson:"category_type"`
	CategoryName string        `bson:"category_name"`
	Question     string        `bson:"question"`
}

type OddOneOutImage struct {
	Id           bson.ObjectId `bson:"_id"`
	CategoryType string        `bson:"category_type"`
	CategoryName string        `bson:"category_name"`
	Filename     string        `bson:"filename"`
}
