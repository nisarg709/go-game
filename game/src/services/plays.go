package services

import (
	"eman/passport/game/src/models"
	"github.com/globalsign/mgo/bson"
	"time"
	"eman/passport/game/src/container"
)

func CreateNewPlay(container *container.Container, playerId string, game string, difficulty int) (*models.Play, error) {

	play := &models.Play{
		Id:         bson.NewObjectId(),
		Player:     playerId,
		Type:       game,
		Difficulty: difficulty,
		StartedAt:  time.Now(),
	}

	if err := container.PlaysRepository.Add(play); err != nil {
		return nil, err;
	}

	return play, nil
}
