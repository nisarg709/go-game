package repositories

import (
	"eman/passport/game/src/models"
	"github.com/globalsign/mgo/bson"
	"time"
)

type PlaysRepository interface {
	Find(id string) (*models.Play, error)
	Add(play *models.Play) error
	Help(play *models.Play) error
	Resume(play *models.Play) error
	Complete(play *models.Play, status bool, score int) error
	CountToday(game string, level int, user string) (int, error)
	Setup(play *models.Play, data interface{}) error
	UpdateAskedQuestionForRetry(play *models.Play, question *models.Question) error
	UpdateAskedQuestionAnswer(play *models.Play, question *models.Question, answer string) error
	UpdateAskedQuestionRemoved(play *models.Play, question *models.Question, removed []bson.ObjectId) error
}

type MongoPlaysRepository struct {
	*MongoRepository
}

func (r *MongoPlaysRepository) Find(id string) (*models.Play, error) {

	play := &models.Play{
		Id: bson.ObjectIdHex(id),
	}

	err := r.Mongo.DB("passport").C("plays").FindId(play.Id).One(play)

	if err != nil {
		return nil, err
	}

	return play, nil
}

func (r *MongoPlaysRepository) Complete(play *models.Play, status bool, score int) error {

	return r.Mongo.DB("passport").C("plays").
		UpdateId(play.Id, bson.M{"$set": bson.M{
			"completed_at": time.Now(),
			"status":       status,
			"score":        score,
		}});
}

func (r *MongoPlaysRepository) Add(play *models.Play) error {
	return r.Mongo.DB("passport").C("plays").Insert(play)
}

func (r *MongoPlaysRepository) Help(play *models.Play) error {

	return r.Mongo.DB("passport").C("plays").
		UpdateId(play.Id, bson.M{"$push": bson.M{"helps": time.Now()}});
}

func (r *MongoPlaysRepository) Resume(play *models.Play) error {

	return r.Mongo.DB("passport").C("plays").
		UpdateId(play.Id, bson.M{"$push": bson.M{"resumes": time.Now()}});
}

func (r *MongoPlaysRepository) Setup(play *models.Play, data interface{}) error {

	return r.Mongo.DB("passport").C("plays").
		UpdateId(play.Id, bson.M{"$set": bson.M{"setup": data}});
}

func (r *MongoPlaysRepository) CountToday(game string, level int, user string) (int, error) {

	query := bson.M{
		"player_id":  user,
		"type":       game,
		"difficulty": level,
		"started_at": bson.M{
			"$gte": time.Now().Truncate(24 * time.Hour),
		},
	}
	return r.Mongo.DB("passport").C("plays").Find(query).Count()
}

func (self *MongoPlaysRepository) UpdateAskedQuestionForRetry(play *models.Play, question *models.Question) error {

	return self.Mongo.DB("passport").C("plays").
		Update(
			bson.M{
				"_id":       play.Id,
				"setup._id": question.Id,
			},
			bson.M{
				"$unset": bson.M{
					"setup.$.meta.answered": "",
				},
				"$set": bson.M{
					"setup.$.meta.removed": question.Meta.Removed,
				},
			},
		)

}

func (self *MongoPlaysRepository) UpdateAskedQuestionAnswer(play *models.Play, question *models.Question, answer string) error {

	return self.Mongo.DB("passport").C("plays").
		Update(
			bson.M{
				"_id":       play.Id,
				"setup._id": question.Id,
			},
			bson.M{
				"$set": bson.M{"setup.$.meta.answered": bson.ObjectIdHex(answer)},
			},
		)

}

func (self *MongoPlaysRepository) UpdateAskedQuestionRemoved(play *models.Play, question *models.Question, removed []bson.ObjectId) error {

	return self.Mongo.DB("passport").C("plays").
		Update(
			bson.M{
				"_id":       play.Id,
				"setup._id": question.Id,
			},
			bson.M{
				"$set": bson.M{"setup.$.meta.removed": removed},
			},
		)
}
