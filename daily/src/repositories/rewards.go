package repositories

import (
	"eman/passport/daily/src/models"
	"github.com/globalsign/mgo/bson"
	"time"
)

type RewardsRepository struct {
	*MongoRepository
}

func (self *RewardsRepository) GiveTickets(userId string, qty int) error {
	return self.incrementTickets(userId, qty);
}

func (self *RewardsRepository) GiveTokens(userId string, qty int) error {
	self.addReward("token", qty, userId, "streak")
	return self.incrementTokens(userId, qty);
}

func (self *RewardsRepository) incrementTickets(userId string, qty int) error {
	self.addTicket(userId, qty)
	return self.Mongo.DB("passport").C("users").
		UpdateId(bson.ObjectIdHex(userId), bson.M{"$inc": bson.M{"tickets": qty}});
}

func (self *RewardsRepository) incrementTokens(userId string, qty int) error {
	return self.Mongo.DB("passport").C("users").
		UpdateId(bson.ObjectIdHex(userId), bson.M{"$inc": bson.M{"tokens": qty}});
}

func (self *RewardsRepository) addReward(rewardType string, rewardQty int, userId string, reason string) error {

	return self.Mongo.DB("passport").C("rewards").Insert(bson.M{
		"action":     "give",
		"user_id":    bson.ObjectIdHex(userId),
		"type":       rewardType,
		"qty":        rewardQty,
		"reason":     reason,
		"created_at": time.Now(),
	})
}

func (self *RewardsRepository) addTicket(userId string, qty int) error {

	for i := 0; i < qty; i++ {
		err := self.Mongo.DB("passport").C("tickets").Insert(models.Ticket{
			UserId:    bson.ObjectIdHex(userId),
			CreatedAt: time.Now(),
		})

		if err != nil {
			return err
		}
	}

	return nil;
}

