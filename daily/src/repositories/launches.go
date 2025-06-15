package repositories

import (
	"eman/passport/daily/src/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type LaunchesRepository struct {
	*MongoRepository
}

func (self *LaunchesRepository) AddLaunch(userId string, consecutivePeriods int) (*models.Launch, error) {

	record := &models.Launch{
		UserId:             userId,
		ConsecutivePeriods: consecutivePeriods,
		CreatedAt:          time.Now(),
		Fresh:              true,
	}

	err := self.Mongo.DB("passport").C("launches").Insert(record)

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (self *LaunchesRepository) GetLastLaunch(userId string) (*models.Launch, error) {

	launch := &models.Launch{}

	err := self.Mongo.DB("passport").C("launches").
		Find(bson.M{"user_id": userId}).
		Sort("-created_at").
		One(launch)

	if err == mgo.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return launch, nil;
}

func (self *LaunchesRepository) CountLaunchesSinceLastStart(userId string, since time.Time) (int, error) {
	return self.Mongo.DB("passport").C("launches").Find(bson.M{
		"user_id": userId,
		"created_at": bson.M{
			"$gt": since,
		},
	}).Count()
}
