package repositories

import (
	"github.com/globalsign/mgo"
)

type MongoRepository struct {
	Mongo *mgo.Session
}
