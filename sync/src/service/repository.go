package service

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"fmt"
	"crypto/md5"
	"encoding/json"
)

type Repository struct {
	db *mgo.Database
}

func Etag(data []byte) (string) {
	etag := fmt.Sprintf("%x", md5.Sum(data))

	return etag
}

func (repository *Repository) GetFreshData() ([]byte, error) {

	continents, err := repository.getContinentsData()

	if err != nil {
		return nil, err
	}

	games, err := repository.getGamesData()
	if err != nil {
		return nil, err
	}

	milestones, err := repository.getMilestonesData()
	if err != nil {
		return nil, err
	}

	data := Response{
		Data{
			Continents: continents,
			Games:      games,
			Progress:   milestones,
		},
	}

	json, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return json, nil
}

func (repository *Repository) getContinentsData() ([]Continent, error) {

	var items []Continent

	query := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "cities",
				"localField":   "_id",
				"foreignField": "continent_id",
				"as":           "cities",
			},
		},
	}

	err := repository.db.C("continents").Pipe(query).All(&items)

	if err != nil {
		return nil, err
	}

	return items, err
}

func (repository *Repository) getGamesData() ([]Game, error) {

	var items []Game

	err := repository.db.C("games").Find(bson.M{}).All(&items)

	if err != nil {
		return nil, err
	}

	return items, err
}

func (repository *Repository) getMilestonesData() ([]Milestone, error) {

	var items []Milestone

	err := repository.db.C("milestones").Find(bson.M{}).All(&items)

	if err != nil {
		return nil, err
	}

	return items, err
}
