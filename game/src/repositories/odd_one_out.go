package repositories

import (
	"eman/passport/game/src/models"
	"github.com/globalsign/mgo/bson"
)

type OddOneOutQuestionRepository struct {
	*MongoRepository
}

type OddOneOutImageRepository struct {
	*MongoRepository
}

func (r OddOneOutImageRepository) GetSimilarRandom(categoryType string, categoryName string, count int) ([]models.OddOneOutImage, error) {
	var questions []models.OddOneOutImage

	pipes := []bson.M{
		{
			"$match": bson.M{
				"category_type": categoryType,
				"category_name": categoryName,
			},
		},
		{
			"$sample": bson.M{
				"size": count,
			},
		},
	}

	err := r.Mongo.DB("passport").C("odd_one_out_images").Pipe(pipes).All(&questions)

	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (r OddOneOutImageRepository) GetDifferentRandom(categoryType string, categoryName string) ([]models.OddOneOutImage, error) {
	var questions []models.OddOneOutImage

	pipes := []bson.M{
		{
			"$match": bson.M{
				"category_type": categoryType,
				"category_name": bson.M{"$ne": categoryName},
			},
		},
		{
			"$sample": bson.M{
				"size": 1,
			},
		},
	}

	err := r.Mongo.DB("passport").C("odd_one_out_images").Pipe(pipes).All(&questions)

	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *OddOneOutQuestionRepository) Random(count int) ([]models.OddOneOutQuestion, error) {

	var questions []models.OddOneOutQuestion

	pipes := []bson.M{
		{
			"$sample": bson.M{
				"size": count,
			},
		},
	}

	err := r.Mongo.DB("passport").C("odd_one_out_questions").Pipe(pipes).All(&questions)

	if err != nil {
		return nil, err
	}

	return questions, nil
}
