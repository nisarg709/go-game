package games

import (
	"eman/passport/game/src/repositories"
	"eman/passport/game/src/models"
	"strconv"
	"errors"
)

func NewGameResolver(repository repositories.GamesRepository, game string, difficulty interface{}) (*GameResolver, error){

	level, err := parseDifficulty(difficulty)

	if err != nil {
		return nil, err
	}

	record, err := repository.FindGame(game, level)

	if err != nil {
		return nil, err
	}

	return &GameResolver{
		record,
		game,
		level,
	}, nil
}

func parseDifficulty(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, errors.New("invalid format")
	}
}


type GameResolver struct {
	Game *models.Game
	GameUid string
	LevelUid int
}

