package repositories

import (
	"eman/passport/game/src/models"
	"github.com/globalsign/mgo/bson"
)

type MilestonesRepository struct {
	*MongoRepository
}

func (r *MilestonesRepository) Achieved(low int, high int) (*models.Milestone, error) {

	milestone := &models.Milestone{}

	query := bson.M{
		"amount": bson.M{
			"$gt":  low,
			"$lte": high,
		},
	}

	err := r.Mongo.DB("passport").C("milestones").Find(query).One(milestone)

	if err != nil {
		return nil, err
	}

	return milestone, nil
}
