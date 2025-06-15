package games

import (
	"eman/passport/game/src/models"
	"eman/passport/game/src/repositories"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/viper"
	"math/rand"
)


func getOddOneOutQuestionsFunc(questionRepo *repositories.OddOneOutQuestionRepository, imageRepo *repositories.OddOneOutImageRepository, level *models.Level) func() ([]*models.Question, error) {

	var s3base = viper.GetString("games.odd_one_out.s3");

	return func() ([]*models.Question, error) {
		var questions []*models.Question

		random, err := questionRepo.Random(level.GetNumberOfQuestions())
		if err != nil {
			return nil, err
		}

		for _, val := range random {

			options, err := imageRepo.GetSimilarRandom(val.CategoryType, val.CategoryName, 5)
			if err != nil {
				return nil, err
			}

			var o []models.QuestionOption
			for _, v := range options {
				o = append(o, models.QuestionOption{
					bson.NewObjectId(),
					s3base + v.Filename,
				})
			}

			answer, err := imageRepo.GetDifferentRandom(val.CategoryType, val.CategoryName)
			if err != nil {
				return nil, err
			}
			a := models.QuestionOption{
				bson.NewObjectId(),
				s3base + answer[0].Filename,
			}

			o = append(o, a)
			rand.Shuffle(len(o), func(i, j int) { o[i], o[j] = o[j], o[i] })

			questions = append(questions, &models.Question{
				Id:       bson.NewObjectId(),
				Question: val.Question,
				Options:  o,
				Answer:   a.Id,
			})
		}



		return questions, nil
	}
}
