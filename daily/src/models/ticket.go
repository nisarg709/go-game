package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Ticket struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	UserId    bson.ObjectId `bson:"user_id"`
	CreatedAt time.Time     `bson:"created_at"`
}
