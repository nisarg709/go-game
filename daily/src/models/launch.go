package models

import (
	"github.com/spf13/viper"
	"time"
)

type Launch struct {
	UserId             string    `bson:"user_id"`
	ConsecutivePeriods int       `bson:"consecutive_periods"`
	CreatedAt          time.Time `bson:"created_at"`
	Fresh              bool      `bson:"-"`
}

func (self *Launch) GetConsecutivePeriods() int {
	return self.ConsecutivePeriods
}

func (self *Launch) IsFromCurrentPeriod() bool {
	period := time.Duration(viper.GetInt("period.hours")) * time.Hour
	check := time.Now().Truncate(period).UTC()
	return self.CreatedAt.After(check)
}

func (self *Launch) IsFromExpiredPeriod() bool {
	period := time.Duration(viper.GetInt("period.hours")) * time.Hour
	check := time.Now().Add(-period).Truncate(period).UTC()
	return self.CreatedAt.Before(check)
}

func (self *Launch) IsFresh() bool {
	return self.Fresh
}

func (self *Launch) NextAvailableAt() int {
	period := time.Duration(viper.GetInt("period.hours")) * time.Hour
	return int(time.Now().Truncate(period).Add(period).Unix())
}
