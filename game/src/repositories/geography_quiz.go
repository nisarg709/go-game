package repositories

import (
	"eman/passport/game/src/models"
	"github.com/globalsign/mgo/bson"
)

type QuizRepository struct {
	*MongoRepository
}

func (r *QuizRepository) Random(count int, difficulty int) ([]*models.Question, error) {

	var questions []*models.Question

	pipes := []bson.M{
		{
			"$match": bson.M{
				"difficulty": recalculateDifficuty(difficulty),
			},
		},
		{
			"$sample": bson.M{
				"size": count,
			},
		},
	}

	err := r.Mongo.DB("passport").C("questions").Pipe(pipes).All(&questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func recalculateDifficuty(i int) int {
	switch i {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9, 10:
		return 3
	case 11, 12, 13, 14:
		return 4
	case 15, 16, 17, 18:
		return 5
	default:
		return 5
	}
}
