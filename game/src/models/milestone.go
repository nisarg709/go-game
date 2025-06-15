package models

import "github.com/globalsign/mgo/bson"

type Milestone struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Amount  int           `json:"amount" bson:"amount"`
	Unlock  string        `json:"unlock" bson:"unlock"`
	Rewards []Reward      `json:"rewards" bson:"rewards"`
}
