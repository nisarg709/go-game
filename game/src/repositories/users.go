package repositories

import (
	"eman/passport/game/src/models"
	"github.com/globalsign/mgo/bson"
	"time"
)

type UsersRepository interface {
	Find(id string) (*models.User, error)
	Charge(user *models.User, amount int, reason string) error
	AddMiles(user *models.User, miles int) error
	AddTicket(user *models.User, tickets int) error
	GetBoosts(id bson.ObjectId) models.BoostCollection
	GetSpentTokens(id bson.ObjectId) (int, error)
	AddBadges(user *models.User, badges []string) error
}

type MongoUsersRepository struct {
	*MongoRepository
}

func (r *MongoUsersRepository) Find(id string) (*models.User, error) {

	user := &models.User{
		Id: bson.ObjectIdHex(id),
	}

	err := r.Mongo.DB("passport").C("users").FindId(user.Id).One(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MongoUsersRepository) AddMiles(user *models.User, miles int) error {

	return r.Mongo.DB("passport").C("users").
		UpdateId(user.Id, bson.M{"$inc": bson.M{"miles": miles}});
}

func (r *MongoUsersRepository) AddTicket(user *models.User, tickets int) error {

	for i := 0; i < tickets; i++ {
		err := r.Mongo.DB("passport").C("tickets").Insert(models.Ticket{
			UserId:    user.Id,
			CreatedAt: time.Now(),
		})

		if err != nil {
			return err
		}
	}

	return r.Mongo.DB("passport").C("users").
		UpdateId(user.Id, bson.M{"$inc": bson.M{"tickets": tickets}});
}

func (r *MongoUsersRepository) Charge(user *models.User, amount int, reason string) error {

	//TODO: more descriptiove reason
	r.Mongo.DB("passport").C("rewards").
		Insert(bson.M{
			"user_id":    user.Id,
			"action":     "charge",
			"type":       "token",
			"qty":        amount,
			"reason":     reason,
			"created_at": time.Now(),
		});
	return r.Mongo.DB("passport").C("users").
		UpdateId(user.Id, bson.M{"$inc": bson.M{"tokens": -amount}});
}

func (r *MongoUsersRepository) GetBoosts(id bson.ObjectId) models.BoostCollection {

	var boosts []models.Boost

	query := bson.M{
		"user_id": id.Hex(),
		"started_at": bson.M{
			"$lte": time.Now(),
		},
		"expires_at": bson.M{
			"$gte": time.Now(),
		},
	}

	err := r.Mongo.DB("passport").C("boosts").Find(query).All(&boosts)

	if err != nil {
		return models.BoostCollection{}
	}

	return models.BoostCollection{boosts}

}

func (r *MongoUsersRepository) AddBadges(user *models.User, badges []string) error {

	return r.Mongo.DB("passport").C("users").
		UpdateId(user.Id, bson.M{"$addToSet": bson.M{"badges": bson.M{"$each": badges}}});
}

func (r *MongoUsersRepository) GetSpentTokens(id bson.ObjectId) (int, error) {

	pipes := []bson.M{
		{
			"$match": bson.M{
				"user_id": id,
				"action": "charge",
			},
		},
		{
			"$group": bson.M{
				"_id": "",
				"total": bson.M{
					"$sum": "$qty",
				},
			},
		},
	}

	var result []bson.M
	err := r.Mongo.DB("passport").C("rewards").Pipe(pipes).All(&result)

	if err != nil {
		return 0, err
	}

	return result[0]["total"].(int), nil

}
