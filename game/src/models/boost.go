package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Boost struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Type      string        `json:"type" bson:"type"`
	UserId    string        `json:"user_id" bson:"user_id"`
	StartsAt  time.Time     `json:"starts_at" bson:"starts_at"`
	ExpiresAt time.Time     `json:"expires_at" bson:"expires_at"`
}

type BoostCollection struct {
	Items []Boost
}

func (self *BoostCollection) IsEmpty() bool {
	if self.Items == nil {
		return true
	}

	return false
}

func (self *BoostCollection) Contains(alias string) bool {
	for _, v := range self.Items {
		if v.Type == alias {
			return true
		}
	}
	return false
}
