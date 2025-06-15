package models

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"math/rand"
	"time"
)

type Question struct {
	Id       bson.ObjectId    `json:"id" bson:"_id" mapstructure:"_id"`
	Question string           `json:"question" bson:"question"`
	Options  []QuestionOption `json:"options" bson:"options"`
	Answer   bson.ObjectId    `json:"-" bson:"answer"`
	Meta     QuestionMeta     `json:"-" bson:"meta,omitempty" mapstructure:"meta"`
}

type QuestionOption struct {
	Id    bson.ObjectId `json:"id" bson:"_id" mapstructure:"_id"`
	Value string        `json:"value" bson:"value"`
}

type QuestionMeta struct {
	Answered bson.ObjectId   `bson:"answered,omitempty"`
	Removed  []bson.ObjectId `bson:"removed"`
}

func NewQuestion() *Question {
	return &Question{Id: bson.NewObjectId()}
}

// CheckAnswer verifies that the passed string is the correct answer
func (self *Question) CheckAnswer(answer string) bool {
	return self.Answer == bson.ObjectIdHex(answer)
}

// IsAnswered verifies if the current question already received and answer by the user
func (self *Question) IsAnswered() bool {
	return self.Meta.Answered != ""
}

// IsCorrect verifies if the question WAS answered correctly
// The answer given by the user is stored in meta.answered
func (self *Question) IsCorrect() bool {
	return self.Meta.Answered == self.Answer
}

// IsHelpAvailable checks if it's possible to use help once more
func (self *Question) IsHelpAvailable() bool {
	return len(self.Meta.Removed) < len(self.Options)-1
}

// IsRemoved checks if a given option was already removed from the options by using Help
// The removed options are stored in meta.removed
func (self *Question) IsRemoved(lookup bson.ObjectId) bool {

	for _, v := range self.Meta.Removed {
		if v == lookup {
			return true
		}
	}
	return false
}

// Reduce is reducing the available options for the user
// All previously removed options + 1 extra are removed from all the possible options
func (self *Question) Reduce() {

	var remaining []QuestionOption
	var candidates []bson.ObjectId

	for _, selected := range self.Options {

		if self.IsRemoved(selected.Id) {
			// if was removed before, don't include it
			continue
		}

		remaining = append(remaining, selected)
		if selected.Id != self.Answer {
			// if it's not the answer add it to the removable candidates
			candidates = append(candidates, selected.Id)
		}

	}

	rand.Seed(time.Now().Unix())
	random := candidates[rand.Intn(len(candidates))];

	for k, option := range remaining {
		if option.Id == random {
			remaining = append(remaining[:k], remaining[k+1:]...)
		}
	}

	self.Meta.Removed = append(self.Meta.Removed, random)
	self.Options = remaining

	fmt.Print("a")
}

func (self *Question) Retry() {
	self.Meta.Removed = append(self.Meta.Removed, self.Meta.Answered)
	self.Meta.Answered = ""
}
