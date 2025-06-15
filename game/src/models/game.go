package models

import "github.com/globalsign/mgo/bson"

type Game struct {
	Uid        string  `json:"uid" bson:"$id"`
	Name       string  `json:"name" bson:"name"`
	HelpCost   int     `json:"help_cost" bson:"help_cost"`
	ResumeCost int     `json:"resume_cost" bson:"resume_cost"`
	ForceCost  int     `json:"force_cost" bson:"force_cost"`
	Levels     []Level `json:"levels" bson:"levels"`
}

type Level struct {
	Difficulty int      `json:"difficulty" bson:"difficulty"`
	Tries      int      `json:"tries" bson:"tries"`
	Rewards    []Reward `json:"rewards" bson:"rewards"`
	Config     bson.M   `json:"config" bson:"config"`
}

type Reward struct {
	Type string `json:"type" bson:"type"`
	Qty  int    `json:"qty" bson:"qty"`
}

type City struct {
	Uid      string    `json:"uid" bson:"uid"`
	Name     string    `json:"name" bson:"name"`
	Requires int       `json:"requires" bson:"requires"`
	Level    CityLevel `json:"level" bson:"level"`
}

type CityLevel struct {
	Difficulty int `json:"difficulty" bson:"difficulty"`
}

func (self *Game) CanPlayMore(plays int) bool {

	level := self.GetLevel()

	if level == nil {
		return false
	}

	return level.Tries > plays
}

func (self *Game) RemainingPlays(plays int) int {

	level := self.GetLevel()

	if level == nil {
		return 0
	}

	return level.Tries - plays
}

func (self *Game) GetRewardedMiles() int {
	level := self.GetLevel()

	if level == nil {
		return 0
	}

	for _, v := range level.Rewards {
		if v.Type == "miles" {
			return v.Qty
		}
	}

	return 0

}

func (self *Game) GetLevel() *Level {
	return &self.Levels[0]
}

func (self *Level) GetConfiguration(key string) interface{} {
	return self.Config[key]
}

func (self *Level) GetNumberOfQuestions() int {
	count, ok := self.GetConfiguration("NumberOfQuestions").(int)
	if !ok || count <= 0 {
		return 5;
	}
	return count
}
