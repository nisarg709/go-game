package games

import (
	"eman/passport/game/src/container"
	"eman/passport/game/src/models"
	"errors"
)

type GameRules interface {
	CostOfHelp() int
	CostOfResume() int
	CanUseHelp(payload map[string]interface{}) bool
	Resume(play *models.Play, payload map[string]interface{}) error
	Checkpoint(play *models.Play, payload interface{}) (bool, error)
	HelpData(payload map[string]interface{}) (interface{}, error)
	SetupData() (interface{}, error)
}

func NewGameRules(play *models.Play, container *container.Container) (GameRules, error) {

	game, err := container.GamesRepository.FindGame(play.Type, play.Difficulty)

	if err != nil {
		return nil, err
	}

	switch play.Type {
	case
		"flight_path",
		"flappy_plane",
		"get_to_plane",
		"my_suitcase",
		"match_the_objects",
		"match_the_pairs",
		"baggage_handler":

		return &ContinuousGame{game}, nil
	case "geography_quiz":
		return NewQuestionGame(game, play, container.PlaysRepository, getGeographyQuizQuestionsFunc(container.QuizRepository, game.GetLevel())), nil
	case "odd_one_out":
		return NewQuestionGame(game, play, container.PlaysRepository, getOddOneOutQuestionsFunc(container.OddOneOutQuestionRepository, container.OddOneOutImageRepository, game.GetLevel())), nil
	case "what_flag":
		return NewQuestionGame(game, play, container.PlaysRepository, getWhatTheFlagQuestions(game.GetLevel())), nil
	default:
		return nil, errors.New("unknown game")
	}
}
