package repositories

import (
	"eman/passport/daily/src/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type EventsRepository struct {
	*MongoRepository
}

func (self *EventsRepository) GetEvent() (*models.Event, error) {

	event := &models.Event{}

	err := self.Mongo.DB("passport").C("popups").
		Find(bson.M{
			"active": true,
			"starts_at": bson.M{
				"$lte": time.Now(),
			},
			"ends_at": bson.M{
				"$gte": time.Now(),
			},
		}).
		Sort("-_id").
		One(event)

	if err == mgo.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	//event.Id = event.GetIdHex()

	return event, nil;
}
