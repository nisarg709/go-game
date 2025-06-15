package container

import (
	"eman/passport/game/src/repositories"
)

type Container struct {
	PlaysRepository             repositories.PlaysRepository
	GamesRepository             repositories.GamesRepository
	UsersRepository             repositories.UsersRepository
	MilestonesRepository        *repositories.MilestonesRepository
	QuizRepository              *repositories.QuizRepository
	OddOneOutQuestionRepository *repositories.OddOneOutQuestionRepository
	OddOneOutImageRepository    *repositories.OddOneOutImageRepository
}
