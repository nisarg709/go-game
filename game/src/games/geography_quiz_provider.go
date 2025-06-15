package games

import (
	"eman/passport/game/src/models"
	"eman/passport/game/src/repositories"
)

func getGeographyQuizQuestionsFunc(repo *repositories.QuizRepository, level *models.Level) func() ([]*models.Question, error) {

	return func() ([]*models.Question, error) {
		return repo.Random(level.GetNumberOfQuestions(), level.Difficulty)
	}
}
