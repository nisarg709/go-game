package repositories

import (
	"eman/passport/game/src/models"
	"github.com/globalsign/mgo/bson"
)

type GamesRepository interface {
	Exists(alias string) bool
	FindGame(alias string, difficulty int) (*models.Game, error)
	FindCityByDifficulty(level int) (*models.City, error)
}

type MongoGamesRepository struct {
	*MongoRepository
}

func (r *MongoGamesRepository) Exists(alias string) bool {

	count, err := r.Mongo.DB("passport").C("games").Find(bson.M{"_id": alias}).Count()

	if err != nil {
		return false
	}

	return count > 0
}

func (r *MongoGamesRepository) FindGame(alias string, difficulty int) (*models.Game, error) {

	game := &models.Game{}

	query := bson.M{"_id": alias, "levels.difficulty": difficulty}
	selects := bson.M{"_id": 1, "name": 1, "help_cost": 1, "resume_cost": 1, "force_cost": 1, "levels.$": 1}

	err := r.Mongo.DB("passport").C("games").Find(query).Select(selects).One(game)

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (r *MongoGamesRepository) FindCityByDifficulty(level int) (*models.City, error) {

	city := &models.City{}

	err := r.Mongo.DB("passport").C("cities").Find(bson.M{"level.difficulty": level}).One(city)

	if err != nil {
		return nil, err
	}

	return city, nil
}
