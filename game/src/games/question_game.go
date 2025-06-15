package games

import (
	"eman/passport/game/src/api"
	"eman/passport/game/src/models"
	"eman/passport/game/src/repositories"
	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/mapstructure"
)

type QuestionGame struct {
	game                 *models.Game
	play                 *models.Play
	repo                 repositories.PlaysRepository
	provideQuestionsFunc func() ([]*models.Question, error)
}

func NewQuestionGame(game *models.Game, play *models.Play, repo repositories.PlaysRepository, provideQuestionsFunc func() ([]*models.Question, error)) *QuestionGame {
	return &QuestionGame{
		game,
		play,
		repo,
		provideQuestionsFunc,
	}
}

// CostOfHelp return the number of tokens needed to continue playin
func (self *QuestionGame) CostOfResume() int {
	return self.game.ResumeCost;
}

// CostOfHelp return the number of tokens to needed to receive question help
func (self *QuestionGame) CostOfHelp() int {
	return self.game.HelpCost;
}

// CanUseHelp checks if the user can use help for the current question
func (self *QuestionGame) CanUseHelp(payload map[string]interface{}) bool {
	//q, err := self.getCurrentQuestion()
	q, err := self.getQuestionById(payload["question_id"].(string))
	if err != nil {
		return false
	}

	return q.IsHelpAvailable()
}

func (self *QuestionGame) Checkpoint(play *models.Play, payload interface{}) (bool, error) {

	//@TODO: this method is returning false instead of an error and this can hide some bugs and errors

	p := payload.(map[string]interface{})
	if _, exists := p["question_id"]; !exists {
		return false, api.NewError(12100, "Missing question_id")
	}
	if _, exists := p["question_answer"]; !exists {
		return false, api.NewError(12100, "Missing question_answer")
	}

	question, err := self.getQuestionById(p["question_id"].(string))
	if err != nil {
		return false, api.NewError(10404, "Not Found")
	}

	if question.IsAnswered() {
		return false, api.NewError(10404, "Question is already answered")
	}

	if err = self.repo.UpdateAskedQuestionAnswer(play, question, p["question_answer"].(string)); err != nil {
		return false, err
	}

	if !question.CheckAnswer(p["question_answer"].(string)) {
		return false, nil
	}

	return true, nil
}

// Resume resets the question and allows the user to continue playing
func (self *QuestionGame) Resume(play *models.Play, payload map[string]interface{}) error {

	//q, err := self.getLastQuestion()
	q, err := self.getQuestionById(payload["question_id"].(string))
	if err != nil {
		return err
	}

	if q.IsAnswered() {
		q.Retry()
		if err = self.repo.UpdateAskedQuestionForRetry(self.play, q); err != nil {
			return nil
		}
	}

	return nil
}

// HelpData prepares and returns the help data when requested by the user
func (self *QuestionGame) HelpData( payload map[string]interface{}) (interface{}, error) {
	//q, err := self.getCurrentQuestion()
	q, err := self.getQuestionById(payload["question_id"].(string))
	if err != nil {
		return nil, err
	}

	q.Reduce()

	if err = self.repo.UpdateAskedQuestionRemoved(self.play, q, q.Meta.Removed); err != nil {
		return nil, err
	}

	return q, nil
}

// SetupData prepares and returns the setup data for the play
func (self *QuestionGame) SetupData() (interface{}, error) {
	if len(self.play.Setup) > 0 {
		return self.oldSetupData()
	}

	return self.freshSetupData()
}

// oldSetupData returns the previously generated Setup data from the database
func (self *QuestionGame) oldSetupData() ([]*models.Question, error) {

	var setup []*models.Question

	for _, el := range self.play.Setup {
		q := &models.Question{}
		if err := mapstructure.Decode(el, &q); err != nil {
			return nil, err
		}
		setup = append(setup, q)
	}

	return setup, nil
}

// freshSetupData prepares a new set of Setup data for the play
func (self *QuestionGame) freshSetupData() ([]*models.Question, error) {

	setup, err := self.provideQuestionsFunc()
	if err != nil {
		return nil, err
	}

	err = self.repo.Setup(self.play, setup)
	if err != nil {
		return nil, err
	}

	return setup, nil
}

// getCurrentQuestion finds the current question and returns it
// The function goes trough all the question for this play and finds the first that is not answered
func (self *QuestionGame) getCurrentQuestion() (*models.Question, error) {

	var q models.Question
	for _, el := range self.play.Setup {
		q = models.Question{}
		if err := mapstructure.Decode(el, &q); err != nil {
			return nil, err
		}
		if !q.IsAnswered() {
			break
		}
	}

	return &q, nil
}


func (self *QuestionGame) getLastQuestion() (*models.Question, error) {

	l := &models.Question{}

	for _, el := range self.play.Setup {
		c := models.Question{}
		if err := mapstructure.Decode(el, &c); err != nil {
			return nil, err
		}
		if !c.IsAnswered() {
			break
		}
		l = &c
	}

	return l, nil
}

// getQuestionById finds a question from the current play by it's ID
func (self *QuestionGame) getQuestionById(questionId string) (*models.Question, error) {

	var q *models.Question

	for _, el := range self.play.Setup {
		temp := &models.Question{}
		if err := mapstructure.Decode(el, temp); err != nil {
			return nil, err
		}
		if temp.Id == bson.ObjectIdHex(questionId) {
			q = temp
			break
		}
	}

	if (q == nil) {
		return nil, api.NewError(12100, "Question not found")
	}

	return q, nil
}
