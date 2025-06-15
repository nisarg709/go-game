package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Play struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Player      string        `json:"-" bson:"player_id,omitempty"`
	Type        string        `json:"type" bson:"type,omitempty"`
	Helps       []interface{} `json:"helps" bson:"helps,omitempty"`
	Resumes     []interface{} `json:"resumes" bson:"resumes,omitempty"`
	Setup       []interface{} `json:"setup" bson:"setup,omitempty"`
	Difficulty  int           `json:"difficulty" bson:"difficulty,omitempty"`
	StartedAt   time.Time     `json:"-" bson:"started_at,omitempty"`
	CompletedAt time.Time     `json:"-" bson:"completed_at,omitempty"`
}

// BelongsTo check if the play belongs to a user with the passed id
func (self *Play) BelongsTo(userId string) bool {
	return self.Player == userId
}

// IsCompleted checks if this play was already completed
func (self *Play) IsCompleted() bool {
	return !self.CompletedAt.IsZero()
}

// HelpsCount counts the number of helps used
func (self *Play) HelpsCount() int {
	return len(self.Helps) + len(self.Resumes)
}
